package main

import "fmt"

// GetMessage returns a greeting message.
func GetMessage(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	message := GetMessage("Nishchal")
	fmt.Println(message)
}
