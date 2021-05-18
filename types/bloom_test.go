package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBloom_MarshalTo(t *testing.T) {
	bloom0 := Bloom{0, 1, 1}

	data, err := bloom0.MarshalJSON()
	require.Nil(t, err)

	var bloom1 Bloom

	err = bloom1.UnmarshalJSON(data)
	require.Nil(t, err)

	require.Equal(t, bloom0, bloom1)
}
