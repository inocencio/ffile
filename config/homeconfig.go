package config

import (
	"fmt"
	"github.com/inocencio/ffile/systemapp"
)

func SetupConfigFiles() {
	var p = systemapp.GetSystemConfigDir()

	fmt.Println(p)
}
