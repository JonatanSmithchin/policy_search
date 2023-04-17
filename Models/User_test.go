package Models

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser_CreateUser(t *testing.T) {
	user := &User{
		UserName: "gx",
		UserPwd:  "123456",
	}
	err := user.CreateUser()
	if err != nil {
		return
	}
	require.NoError(t, err)
}

func TestUser_ChangeAge(t *testing.T) {
	user := &User{
		UserID: 1,
	}
	err := user.ChangeAge(18)
	require.NoError(t, err)
}

func TestUser_ChangeEmail(t *testing.T) {
	user := &User{UserID: 1}
	err := user.ChangeEmail("lxx18912405977@outlook.com")
	require.NoError(t, err)
}

func TestUser_ChangePassword(t *testing.T) {
	user := &User{
		UserID:  1,
		UserPwd: "44913730d",
	}
	err := user.ChangePassword(user.UserPwd)
	require.NoError(t, err)
}

func TestUser_DeleteUser(t *testing.T) {
	user := &User{UserID: 2}
	err := user.DeleteUser()
	require.NoError(t, err)
}
