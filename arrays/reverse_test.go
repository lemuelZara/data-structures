package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{
			name:     "should reverse the 'Hi My name is Lemuel' string",
			param:    "Hi My name is Lemuel",
			expected: "leumeL si eman yM iH",
		},
		{
			name:     "should reverse the 'apwoekjd' string",
			param:    "apwoekjd",
			expected: "djkeowpa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.param)

			require.Equal(t, tt.expected, result)
		})
	}
}
