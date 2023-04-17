package Controllers_test

import (
	"log"
	"police_search/Controllers"
	"testing"
)

func TestCasbinList(t *testing.T) {
	auths := Controllers.CasbinList("admin")
	log.Println(auths)
}
