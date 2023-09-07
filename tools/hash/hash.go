package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

type Hash interface {
	MD5(input string) (string, error)
	SHA1(input string) (string, error)
	SHA256(input string) (string, error)
	SHA512(input string) (string, error)
}

type HashImpl struct {
}

func (obj *HashImpl) MD5(input string) (string, error) {
	hash := md5.New()
	hash.Write([]byte(input))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue), nil
}

func (obj *HashImpl) SHA1(input string) (string, error) {
	hash := sha1.New()
	hash.Write([]byte(input))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue), nil
}

func (obj *HashImpl) SHA256(input string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue), nil
}

func (obj *HashImpl) SHA512(input string) (string, error) {
	hash := sha512.New()
	hash.Write([]byte(input))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue), nil
}
