package epochid

import "slices"

type Epochs []EpochID

func (e Epochs) AsString() []string {
	result := make([]string, len(e))

	for ix, epoch := range e {
		result[ix] = epoch.String()
	}

	return result
}

// Contains would come handy in tests.
func (e Epochs) Contains(item int) bool {
	return slices.Contains(
		e,
		EpochID(item),
	)
}
