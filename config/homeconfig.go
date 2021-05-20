package config

import (
	"fmt"
	"github.com/inocencio/ffile/systemapp"
)

func SetupConfigFiles() {
	var p = systemapp.ConfigDir()

	fmt.Println(p)
}
