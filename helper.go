package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type Password struct {
	Salt     [4]byte
	Password string
}

func (p *Password) generateSalt() {
	rand.Seed(time.Now().UTC().UnixNano())

	// 32-bit random salt
	p.Salt[0] = byte(rand.Int31() * 255)
	p.Salt[1] = byte(rand.Int31() * 255)
	p.Salt[2] = byte(rand.Int31() * 255)
	p.Salt[3] = byte(rand.Int31() * 255)
}

// See https://rabbitmq.com/passwords.html#computing-password-hash
func (p *Password) Hash() []byte {
	if p.Salt[0] == 0 && p.Salt[1] == 0 && p.Salt[2] == 0 && p.Salt[3] == 0 {
		p.generateSalt()
	}

	// password in bytes
	pwdBytes := []byte(p.Password)
	// append salt to the start of pwdBytes
	pwdBytes = append(p.Salt[:], pwdBytes...)

	// hash pwdBytes
	pwdHashed := sha256.Sum256(pwdBytes)

	// append the salt again
	var final []byte
	final = pwdHashed[:]
	final = append(p.Salt[:], final...)

	return final
}

func (p *Password) HashBase64() string {
	hash := p.Hash()
	return base64.StdEncoding.EncodeToString(hash[:])
}

func (p *Password) HashString() string {
	return fmt.Sprintf("%x", p.Hash())
}

func hashPassword(password string) []byte {
	p := &Password{
		Password: password,
	}

	return p.Hash()
}

func hashPasswordString(password string) string {
	p := &Password{
		Password: password,
	}

	return p.HashString()
}

func hashPasswordBase64(password string) string {
	p := &Password{
		Password: password,
	}

	return p.HashBase64()
}

// normalizing names such as vhosts, queues, exchanges, etc.
// before put to request
// traper@dev-id/air/humid -> traper@dev-id%2fair%2fhumid
func normalizeNames(name string) string {
	return url.QueryEscape(name)
}
