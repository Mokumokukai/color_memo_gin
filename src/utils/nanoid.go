package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var alphanum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func AlphaNumNanoID(len int) (string, error) {
	return gonanoid.Generate(alphanum, len)
}
