//go:generate go run github.com/ImGajeed76/charmer/tools/generate

package main

import (
	"tfutils-go/internal/registry"
	"github.com/ImGajeed76/charmer/pkg/charmer"
)

func main() {
	charmer.Run(registry.RegisteredCharms)
}