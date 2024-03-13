package funcions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			name: "hello",
			in:   []byte("hello"),
			out:  []byte("olleh"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.in)
			require.Equal(t, tt.in, tt.out)
		})
	}
}
