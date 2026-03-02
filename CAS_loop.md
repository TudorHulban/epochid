# CAS loop considerations

## Why the CAS Loop is Needed in `getSequenceID()`

The incremental mode of `epochid` must guarantee **unique differentiators** (0–9999) within each 100 ms time window — even when many goroutines call `getSequenceID()` simultaneously.

Without proper synchronization, a naive implementation would suffer from **lost updates** and **duplicate IDs**.

Broken naive version (do NOT use):

```go
current := gen.sequenceID.Load()

next := current + 1
if next >= _sequenceLimit {
    next = 0
}

gen.sequenceID.Store(next)

return gen.precomputedIDs[next]
```

What goes wrong under contention

Imagine _sequenceLimit = 10000 and three goroutines that all read current as 9998 because they issue the calls at same time.  
All three compute next as 9999.  
Goroutine A stores 9999 → `returns precomputed[9999]`  
Goroutine B stores 9999 (overwrites, no change) → `returns precomputed[9999]`  
Goroutine C stores 9999 → returns `precomputed[9999]`  

Result: three goroutines receive the exact same ID, violating uniqueness.  
This is the classic lost update problem — multiple increments are lost because the read → compute → write sequence is not atomic.  

## How the CAS loop achieves correctness

The loop implements optimistic concurrency control using a single atomic Compare-And-Swap operation:

a. Read the current value (Load)  
b. Compute the desired next value (current + 1, with wrap-around to 0)  
c. Attempt to atomically update only if the value is still exactly what we read (CompareAndSwap).  
Success → only goroutine that advanced the counter at the moment → safe to return the corresponding precomputed string.  
Failure → another goroutine won the race and already changed the value → retry from the beginning.  

### Key Properties Achieved by the CAS Loop

| Property                          | Description                                                                 | How the CAS Loop Guarantees It                                                                 |
|-----------------------------------|-----------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| **Uniqueness**                    | No two goroutines receive the same sequence number in the same time window  | Only one goroutine succeeds in the CAS per logical increment; others retry                      |
| **Atomic increment + wrap-around**| Increment and rollover (9999 → 0) happen as a single indivisible operation  | The wrap decision is computed and validated inside the CAS attempt — no separate race window   |
| **No overshooting / bounds safety**| Never produces a value ≥ `_sequenceLimit`                                  | Invalid `next` values (≥ limit) are never stored; CAS only succeeds with correct wrapped value |
| **Lock-free hot path**            | No mutex or blocking synchronization is used                                | Relies exclusively on atomic `Load` + `CompareAndSwap` operations                              |
| **Eliminates TOCTOU race**        | Closes the Time-Of-Check To Time-Of-Use gap                                 | The check (current value) and the update are performed atomically in one CAS instruction       |
| **Progress guarantee**            | Every waiting goroutine eventually succeeds                                 | Under contention, retries are short and fair; no starvation or livelock in practice            |
| **Minimal overhead in common case**| Very low latency when contention is low or absent                         | Most calls succeed on the first CAS attempt (~few nanoseconds overhead)                        |
| **Correctness under high contention** | Safe even when dozens/hundreds of goroutines call simultaneously         | Retries handle real races efficiently; system makes forward progress                           |


## Performance in practice

Low contention (most real workloads): CAS succeeds on first try → negligible overhead (~few nanoseconds).  
High contention (many goroutines at once): some retry 2–10× (still very fast — just atomic instructions).  
No artificial delay or backoff — retries are purely to resolve real races.