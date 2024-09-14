package epochid

type Epochs []EpochID

func (e Epochs) AsString() []string {
	result := make([]string, len(e))

	for ix, epoch := range e {
		result[ix] = epoch.String()
	}

	return result
}
