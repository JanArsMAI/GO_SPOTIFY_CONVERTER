package entity

import "time"

type Token struct {
	Token     string        `json:"access_token"`
	Bearer    string        `json:"token_type"`
	ExpiresIn time.Duration `json:"expires_in"`
}
