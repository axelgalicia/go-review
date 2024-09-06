package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	messages := []string{
		"Hello, %v. Welcome",
		"Did I say hi to you already %v!",
		"Welcome back %v",
	}

	return messages[rand.Intn(len(messages))]
}
