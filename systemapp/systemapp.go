package systemapp

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func CheckErr(err error, execFatal bool) {
	if err != nil {
		fmt.Println(err)

		if execFatal {
			os.Exit(-1)
		}
	}
}

func GetSystemConfigDir() string {
	var p, err = os.UserConfigDir()
	CheckErr(err, true)

	return p
}

func GetSystemHomeDir() string {
	var p, err = os.UserHomeDir()
	CheckErr(err, true)

	return p
}

func GetSystemCacheDir() string {
	var p, err = os.UserCacheDir()
	CheckErr(err, true)

	return p
}

func CloseFile(f *os.File, ferror bool) {
	var err = f.Close()

	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Error on closing file: %v\n", err)

		if ferror {
			os.Exit(-1)
		}
	}
}

func FNormalizePath(ffpath string) string {
	return strings.ReplaceAll(strings.TrimSpace(ffpath), "\\", "/")
}

func NormalizePath(fp string, fname string) string {
	return FNormalizePath(path.Join(fp, fname))
}

func FReadFileBytes(ffpath string, ferror bool) *[]byte {
	var fp, fname = filepath.Split(ffpath)
	return ReadFileBytes(fp, fname, ferror)
}

func ReadFileBytes(fp string, fname string, ferror bool) *[]byte {
	fp = FNormalizePath(fp)
	var b, err = ioutil.ReadFile(path.Join(fp, fname))
	CheckErr(err, ferror)

	return &b
}

func FReadFile(ffpath string, ferror bool) *string {
	var fp, fname = filepath.Split(ffpath)
	return ReadFile(fp, fname, ferror)
}

func ReadFile(fp string, fname string, ferror bool) *string {
	var b = ReadFileBytes(fp, fname, ferror)
	var s = string(*b)
	return &s
}

func FReadFileScanner(ffpath string, ferror bool) (*bufio.Scanner, *os.File) {
	var fp, fname = filepath.Split(ffpath)
	return ReadFileScanner(fp, fname, ferror)
}

func ReadFileScanner(fp string, fname string, ferror bool) (*bufio.Scanner, *os.File) {
	fp = FNormalizePath(fp)
	var f, err = os.Open(path.Join(fp, fname))
	CheckErr(err, ferror)

	return bufio.NewScanner(f), f
}

func FWriteFile(ffpath string, content *string, ferror bool) {
	var fp, fname = filepath.Split(ffpath)
	WriteFile(fp, fname, content, ferror)
}

func WriteFile(fp string, fname string, content *string, ferror bool) {
	var fullpath = NormalizePath(fp, fname)
	var err = ioutil.WriteFile(fullpath, []byte(*content), 0644)
	CheckErr(err, ferror)
}

func WriteSliceSlice(fp string, fname string, content *[]string, ferror bool) {
	var fullpath = NormalizePath(fp, fname)

	f, err := os.Create(fullpath)
	CheckErr(err, ferror)
	defer CloseFile(f, ferror)

	for _, line := range *content {
		_, err := f.WriteString(line)
		CheckErr(err, ferror)
	}
}

func OpenDefault(ffpath string) error {
	var err error

	if _, err = os.Stat(ffpath); os.IsNotExist(err) {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", ffpath).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", ffpath).Start()
	case "darwin":
		err = exec.Command("open", ffpath).Start()
	default:
		err = fmt.Errorf("Unsupported platform\n")
	}

	if err != nil {
		return err
	}

	return nil
}
