package warppipe

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {
	w, err := NewWriter[string]("/tmp/warppipe.test.writer.queue")
	require.NoError(t, err)

	// start reader to unblock writes
	go os.ReadFile("/tmp/warppipe.test.writer.queue")

	// send
	send := []string{"hello", "mr blobby", "goodbye"}
	for _, v := range send {
		require.NoError(t, w.Write(v))
	}
	require.NoError(t, w.Close())
}
