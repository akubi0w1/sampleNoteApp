package auth

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	const pw = "password"
	hash, err := PasswordHash(pw)
	if err != nil {
		t.Error("fail password hash")
	}

	err = PasswordVerify(hash, pw)
	if err != nil {
		t.Error("fail password verify")
	}

	err = PasswordVerify(hash, "pass")
	if err == nil {
		t.Error(fmt.Sprintf("fail password verify. hash=%s, truePassword=%s. password=pass", hash, pw))
	}

}
