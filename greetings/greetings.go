package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello %v, Welcome back!", name)
	return message
}
