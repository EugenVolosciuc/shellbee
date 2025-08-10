package storage

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
)

func getStorageFileLocation() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Error reading the user's home directory: %w", homeDir)
	}

	fileLocation := path.Join("shellbee", "store.json")

	var finalFileLocation string

	switch runtime.GOOS {
	case "windows/386":
		fallthrough
	case "windows/amd64":
		fallthrough
	case "windows/arm":
		fallthrough
	case "windows/arm64":
		finalFileLocation = path.Join(homeDir, fileLocation)
	case "darwin/amd64":
		fallthrough
	case "darwin/arm64":
		finalFileLocation = path.Join(homeDir, "Library", "Application Support", fileLocation)
	case "linux/386":
		fallthrough
	case "linux/amd64":
		fallthrough
	case "linux/arm":
		fallthrough
	case "linux/arm64":
		fallthrough
	case "linux/loong64":
		fallthrough
	case "linux/mips":
		fallthrough
	case "linux/mips64":
		fallthrough
	case "linux/mips64le":
		fallthrough
	case "linux/mipsle":
		fallthrough
	case "linux/ppc64":
		fallthrough
	case "linux/ppc64le":
		fallthrough
	case "linux/riscv64":
		fallthrough
	case "linux/s390x":
		xdgDataHomeVar, found := os.LookupEnv("XDG_DATA_HOME")

		if found {
			finalFileLocation = path.Join(xdgDataHomeVar, fileLocation)
		} else {
			finalFileLocation = path.Join(homeDir, ".local", "share", fileLocation)
		}
	default:
		return "", errors.New("unsupported OS")
	}

	return finalFileLocation, nil
}

func readStorageFile() error {
	return nil
}

func ReadKey() {

}

func ListKeys() {

}
