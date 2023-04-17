package Controllers

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetTendency(t *testing.T) {
	tendencies := []string{
		"pc", "history", "agriculture",
	}
	err := SetTendency(1, tendencies)
	require.NoError(t, err)
}

func TestRemoveTendency(t *testing.T) {
	tendencies := []string{
		"pc", "history", "agriculture",
	}
	err := RemoveTendency(1, tendencies)
	require.NoError(t, err)
}

func TestGetTendency(t *testing.T) {
	tendencies, err := GetTendency(1)
	require.NoError(t, err)
	for _, tendency := range tendencies {
		fmt.Println(tendency)
	}
}
