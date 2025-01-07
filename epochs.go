package epochid

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
	for _, epoch := range e {
		if epoch == EpochID(item) {
			return true
		}
	}

	return false
}
