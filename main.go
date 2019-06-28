package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Printf("Prism Installer for %s (%s)\n\n", runtime.GOOS, runtime.GOARCH)

	fmt.Println("Installing...")
	if err := DownloadFile(); err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}

func DownloadFile() error {
	var file string
	url := "https://github.com/PrismLang/binaries/raw/master/prism-" + runtime.GOOS + "-" + runtime.GOARCH

	switch runtime.GOOS {
	case "windows":
		url += ".exe"

		installDir := os.Getenv("SYSTEMROOT")
		file = filepath.Join(installDir, "prism.exe")

	default:
		installDir := "/usr/local/bin"
		file = filepath.Join(installDir, "prism")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
