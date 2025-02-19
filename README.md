# Epoch ID

## Unicity among containers

What should be the minimum for creating an ID that would have chances of being unique among containers?
And what is important within a distributed system for an ID?  
What about the below:  
a. ID is of type `uint64` with maximum being `18446744073709551615` (20 digits).  
b. epoch time  
c. container ID
d. some random number  
e. some sequence

Why not process ID? Due to the fact that a container should only run one process.

## Implementation - Epoch time

A value up to a 1/10 of a second was chosen for good precision.  
So `1654589015` is seconds, length is 10 digits. Our ID would use 11 digits.  
This gives us another 9 positions.

## Implementation - Container ID

The value would be read from `/etc/machine-id`.
Would cover 5 postions.

## Implementation - Random number

Would cover 4 positions.  
The value would be compensated with the epoch time nano second positions to 
always reach the length.
The constructor to be used is `NewIDRandom`.

## Implementation - Sequence number

Would cover 4 positions as it replaces the random number in `NewIDIncremental10K`.  
This constructor should be used with traffic up to 10000 requests per 1/10 of a second. It is faster than the random one as per the benchmarks and was tested for race conditions.

## Process ID

Go routines do not spawn new process IDs as per:

```go
package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	res := make(map[int]struct{})
	var wg sync.WaitGroup

	pid := func() {
		p := os.Getpid()

		if _, exists := res[p]; !exists {
			res[p] = struct{}{}
		}

		wg.Done()
	}

	no := 50
	wg.Add(no)

	for i := 0; i < no; i++ {
		go pid()
	}

	wg.Wait()

	fmt.Println(res)
}
```

## Resources

```html
https://www.sobyte.net/post/2022-01/go-snowflakeid/
```
