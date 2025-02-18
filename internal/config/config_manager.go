package config

import (
	"github.com/ImGajeed76/charmer/pkg/charmer/config"
	"github.com/ImGajeed76/charmer/pkg/charmer/console"
	"github.com/ImGajeed76/charmer/pkg/charmer/path"
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

	changeSettings, changeErr := console.YesNo(console.YesNoOptions{
		Prompt:     "Do you want to change the SFTP settings?",
		DefaultYes: false,
		YesText:    "Yes, change settings",
		NoText:     "No, keep existing settings",
	})

	if changeErr != nil {
		panic(changeErr)
	}

	if changeSettings {
		err = globalCfg.Delete("sftp-username")
		if err != nil {
			return
		}
		err = globalCfg.Delete("sftp-password")
		if err != nil {
			return
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
