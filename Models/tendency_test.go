package Models

import (
	"fmt"
	"github.com/stretchr/testify/require"
	Mysql "police_search/Databases"
	"testing"
)

func TestTendency_AddTendency(t *testing.T) {
	tendency := Tendency{
		UserId:      3,
		Description: "ai",
	}
	err := tendency.AddTendency(Mysql.DB)
	require.NoError(t, err)
}

func TestTendency_DeleteTendency(t *testing.T) {
	tendency := Tendency{
		UserId:      1,
		Description: "ai",
	}
	err := tendency.DeleteTendency(Mysql.DB)
	require.NoError(t, err)
}

func TestTendency_FindTendency(t *testing.T) {
	tendencies, err := FindTendency(1)
	require.NoError(t, err)
	for _, tendency := range tendencies {
		fmt.Println(tendency)
	}
}
