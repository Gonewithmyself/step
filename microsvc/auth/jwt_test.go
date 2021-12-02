package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func Test_genToken(t *testing.T) {
	secret = "5"
	sh := sha256.New()
	sh.Write([]byte(secret))
	secret = hex.EncodeToString(sh.Sum(nil))
	tk, er := genToken(&payload{
		Usr: "john",
		Psw: "123456",
	})
	t.Log(tk, er)
	if er != nil {
		return
	}

	pl, er := parseToken(tk)
	t.Log(pl, er)
}
