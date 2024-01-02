package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A Slice of Random Message Formats
	formats := []string{
		"Hi %v, Welcome!",
		"Hail %v, Well met.",
		"Ssup, %v",
	}
	return formats[rand.Intn(len(formats))]
}
