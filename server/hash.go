package main

import (
	"log"

	"github.com/roberthodgen/password"
)

func Decrypt(passwordToHash string) (string, string) {
	pass := password.New()
	if err := pass.Generate(passwordToHash); err != nil {
		log.Fatal("Something went wrong while generating hash for password: ", err)
	}
	return pass.Hash, pass.Salt
}

func Encrypt(passwordAttempt string, hash string, salt string) bool {
	checker := password.NewChecker(hash, salt)
	if err := checker.Check(passwordAttempt); err == password.ErrIncorrect {
		return false
	} else if err != nil {
		return false
	}

	return true
}
