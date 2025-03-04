package config

import (
	"github.com/ImGajeed76/charmer/pkg/charmer/config"
	"github.com/ImGajeed76/charmer/pkg/charmer/console"
	"github.com/ImGajeed76/charmer/pkg/charmer/path"
	"os"
	"sync"
)

var (
	globalCfg *config.Config
	once      sync.Once
)

func InitConfig() {
	once.Do(func() {
		cfg, err := config.New("charmer-testing")
		if err != nil {
			panic(err)
		}
		globalCfg = cfg
	})

	err := globalCfg.SetDefault("sftp-hostname", "sftp.tfbern.ch")
	if err != nil {
		panic(err)
	}

	err = globalCfg.SetDefault("sftp-port", "22")
	if err != nil {
		panic(err)
	}

	// Clear SFTP credentials if requested
	if len(os.Args) > 1 && os.Args[1] == "--clear-sftp" {
		err = globalCfg.Delete("sftp-username")
		if err != nil {
			panic(err)
		}
		err = globalCfg.Delete("sftp-password")
		if err != nil {
			panic(err)
		}
	}

	if !globalCfg.Exists("sftp-username") {
		_, usernameErr := globalCfg.SetFromInput("sftp-username", console.InputOptions{
			Prompt:   "Enter the SFTP username",
			Required: true,
		})
		if usernameErr != nil {
			panic(usernameErr)
		}
	}

	if !globalCfg.Exists("sftp-password") {
		_, passwordErr := globalCfg.SetFromInput("sftp-password", console.InputOptions{
			Prompt: "Enter the SFTP password (will be stored securely)",
		})
		if passwordErr != nil {
			panic(passwordErr)
		}
	}
}

func Cfg() *config.Config {
	return globalCfg
}

func GetSFTPConfig() *path.SFTPConfig {
	return &path.SFTPConfig{
		Host:     globalCfg.Get("sftp-hostname"),
		Port:     globalCfg.Get("sftp-port"),
		Username: globalCfg.Get("sftp-username"),
		Password: globalCfg.Get("sftp-password"),
	}
}
