package models

import "testing"

func TestCanFindUserByEmail(t *testing.T) {
	_, err := GetUserByEmail("john@example.org")

	if err != nil {
		t.Fail()
	}
}
