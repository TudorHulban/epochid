package epochid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadHostID(t *testing.T) {
	idHost, errRead := readContainerID()
	require.NoError(t, errRead)

	value := hash(idHost, 4)

	require.Len(t,
		value,
		4,
	)

	fmt.Println(value)
}
