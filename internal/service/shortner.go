package service

import (
	"errors"
	"net/url"
	"url-shortner/internal/store"
)

type Shortner struct {
	store      store.Store
	codeLength int
	maxRetries int
}

var (
	ErrInvalidURL     = errors.New("Invalid url")
	ErrCodeGeneration = errors.New("Could not generate unique code")
)

func NewShortener(s store.Store) *Shortner {

	return &Shortner{
		store:      s,
		codeLength: 6,
		maxRetries: 5,
	}
}

func (s *Shortner) Shortner(longURL string) (string, error) {

	if !isValidURL(longURL) {
		return "", ErrInvalidURL
	}

	for i := 0; i < s.maxRetries; i++ {

		code, err := GenerateCode(s.codeLength)

		if err != nil {
			return "", err
		}

		if s.store.Exists(code) {
			continue
		}

		if err := s.store.Save(code, longURL); err != nil {
			return "", err
		}

		return code, nil
	}

	return "", ErrCodeGeneration

}

func isValidURL(raw string) bool {

	u, err := url.ParseRequestURI(raw)

	if err != nil {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}
