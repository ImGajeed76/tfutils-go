//go:generate go run github.com/ImGajeed76/charmer/tools/generate

package main

import (
	"github.com/ImGajeed76/charmer/pkg/charmer"
	"tfutils-go/internal/config"
	"tfutils-go/internal/registry"
)

func main() {
	config.InitConfig()

	charmer.Run(registry.RegisteredCharms)
}
