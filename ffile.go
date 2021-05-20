package main

import (
	"github.com/inocencio/ffile/config"
	. "github.com/inocencio/ffile/systemapp"
)

func main() {
	config.SetupConfigFiles()

	var content = "conteudo de teste\narquivo de teste"
	WriteFile(GetSystemHomeDir(), "arquivo_teste.txt", &content, true)

	var sl = []string{"Dados de entrada\n", "Arquivos validos\n", "Projeto teste"}
	WriteFileSlice(GetSystemHomeDir(), "arquivo_slice.txt", &sl, true)
	WriteFileSliceAppend(NormalizePath(GetSystemHomeDir(), "arquivo_append.txt"), &sl, false)
}
