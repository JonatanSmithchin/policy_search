package Controllers

import (
	"fmt"
	"testing"
)

func TestGetAllAdmins(t *testing.T) {
	admins := GetAllAdmins()
	for _, admin := range *admins {
		fmt.Println(admin)
	}
}
