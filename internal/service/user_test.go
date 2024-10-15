package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncrypt(t *testing.T) {
	password := "1234#abc"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(encrypted))
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}
