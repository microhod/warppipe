package warppipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	r, err := NewReader[string]("/tmp/warppipe.test.reader.queue")
	require.NoError(t, err)

	// sent
	sent := []string{"hello", "mr blobby", "goodbye"}
	w, err := NewWriter[string]("/tmp/warppipe.test.reader.queue")
	require.NoError(t, err)
	go func() {
		for _, v := range sent {
			require.NoError(t, w.Write(v))
		}
	}()

	// read
	read := make([]string, len(sent))
	for i := range read {
		require.NoError(t, r.Read(&read[i]))
	}
	assert.Equal(t, sent, read)

	//cleanup
	require.NoError(t, w.Close())
	require.NoError(t, r.Close())
}
