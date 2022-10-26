package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	t.Run("should merge sorted array", func(t *testing.T) {
		a := []int{1, 3, 6}
		b := []int{2, 4, 6, 8}

		res := Merge(a, b)

		require.Equal(t, []int{1, 2, 3, 4, 6, 6, 8}, res)
	})
}
