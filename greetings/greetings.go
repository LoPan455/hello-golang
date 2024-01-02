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

// takes in ["ed", "sam", "sally"], returns [ "ed":"Hail, ed, well met.", "sam": "Sssup, sam", etc...]
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
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
