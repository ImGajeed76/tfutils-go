package charms

import (
	"fmt"
	"github.com/ImGajeed76/charmer/pkg/charmer/console"
)

// Greeting godoc
// @Charm
// @Title Greeting
// @Description
// # Greeting
// ## Description
// This is a simple greeting function that asks for a name and greets the user.
func Greeting() {
	name, _ := console.Input(console.InputOptions{
		Prompt: "What is your name?",
	})

	fmt.Printf("Hello, %s!\n", name)
}
