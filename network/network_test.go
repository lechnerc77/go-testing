package network_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/lechnerc77/gotestdemo/network"
	"github.com/stretchr/testify/assert"
)

func TestReadHosts(t *testing.T) {

	filesystem := fstest.MapFS{
		"hosts": {Data: []byte("127.0.0.1 localhost\n")},
	}

	content, err := network.ReadHosts(filesystem)

	assert.Equal(t, "127.0.0.1 localhost\n", content)
	assert.NoError(t, err)
}

func TestReadHostsFromFile(t *testing.T) {

	tempDirectory, error := os.MkdirTemp("", "test-")
	assert.NoError(t, error)
	defer os.RemoveAll(tempDirectory)

	hostfile := filepath.Join(tempDirectory, "hosts")
	error = os.WriteFile(hostfile, []byte("127.0.0.1 localhost\n"), 0644)

	assert.NoError(t, error)

	content, err := network.ReadHostsFromFile(hostfile)
	assert.Equal(t, "127.0.0.1 localhost\n", content)
	assert.NoError(t, err)
}
