package consul

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestFindServer(t *testing.T) {
	host, err := FindServer("test")
	log.Print(host)
	require.NoError(t, err)

}
