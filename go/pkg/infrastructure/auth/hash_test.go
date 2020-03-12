package auth

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		_, err := PasswordHash("password")
		if err != nil {
			t.Error("fail!")
		}
	})
}