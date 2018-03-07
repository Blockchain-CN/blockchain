package models

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	u, err := Login("luda")
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("%v", u)
}
