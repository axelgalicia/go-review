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

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)
	// Lool through the recieved slices of names, calling
	// the Hello function to get the message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
