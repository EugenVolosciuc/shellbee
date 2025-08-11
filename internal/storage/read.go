package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

type Alias string
type AliasesMap map[string]string

type StorageFileData struct {
	Aliases AliasesMap `json:"aliases"`
}

type KeyNotFoundError struct {
	key string
}

func (err *KeyNotFoundError) Error() string {
	return fmt.Sprintf("%s key not found", err.key)
}

func getStorageFileLocation() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Error reading the user's home directory: %w", homeDir)
	}

	fileLocation := path.Join("shellbee", "store.json")

	var finalFileLocation string

	switch runtime.GOOS {
	case "windows":
		fallthrough
	case "windows/386":
		fallthrough
	case "windows/amd64":
		fallthrough
	case "windows/arm":
		fallthrough
	case "windows/arm64":
		finalFileLocation = path.Join(homeDir, fileLocation)
	case "darwin":
		fallthrough
	case "darwin/amd64":
		fallthrough
	case "darwin/arm64":
		finalFileLocation = path.Join(homeDir, "Library", "Application Support", fileLocation)
	case "linux":
		fallthrough
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

func checkStorageFileExists() (bool, error) {
	filePath, err := getStorageFileLocation()

	if err != nil {
		return false, err
	}

	_, err = os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func readStorageFile() (StorageFileData, error) {
	var result StorageFileData

	finalLocation, err := getStorageFileLocation()

	if err != nil {
		return result, fmt.Errorf("could not get storage file location: %w", err)
	}

	jsonFile, err := os.Open(finalLocation)

	if err != nil {
		return result, fmt.Errorf("could not open the storage file: %w", err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &result)

	return result, nil
}

func ReadKey(key string) (string, error) {
	storageFileData, err := readStorageFile()

	if err != nil {
		return "", err
	}

	command, ok := storageFileData.Aliases[key]

	if !ok {
		return "", &KeyNotFoundError{key}
	}

	return command, nil
}

func ListKeys() {

}
