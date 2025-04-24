package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

type PKCE struct {
	CodeVerifier  string
	CodeChallenge string
}

func GeneratePKCE() (*PKCE, error) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}
	verifier := base64.RawURLEncoding.EncodeToString(bytes)

	sha := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sha[:])

	return &PKCE{
		CodeVerifier:  verifier,
		CodeChallenge: challenge,
	}, nil
}
