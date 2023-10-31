package passwd

import "testing"

var obj Password

func init() {
	obj = &PasswordImpl{Length: 15}
}

func TestGeneratePassword_RandomPassword(t *testing.T) {
	output, _ := obj.RandomPassword()
	t.Log(output)
}
