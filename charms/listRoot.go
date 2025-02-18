package charms

import (
	"github.com/ImGajeed76/charmer/pkg/charmer/path"
	"tfutils-go/internal/config"
)

// ListRoot godoc
// @Charm
// @Title ListRoot
// @Description
// # ListRoot
// ## Description
// Lists the root of the sftp server
func ListRoot() {
	rootPath := path.New("/", config.GetSFTPConfig())
	list, err := rootPath.List()
	if err != nil {
		return
	}

	for _, p := range list {
		println(p.String())
	}
}
