package main

import (
	"github.com/inocencio/ffile/config"
	. "github.com/inocencio/ffile/systemapp"
)

func main() {
	config.SetupConfigFiles()

	var content = "conteudo de teste\narquivo de teste"
	WriteFile(GetSystemHomeDir(), "arquivo_teste.txt", &content, true)
}
