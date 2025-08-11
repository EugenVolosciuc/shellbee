package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func createStorageFile() error {
	filePath, err := getStorageFileLocation()

	if err != nil {
		return err
	}

	jsonStr, _ := json.Marshal(StorageFileData{})

	err = os.MkdirAll(filepath.Dir(filePath), 0755)

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, []byte(jsonStr), 0644)

	if err != nil {
		return err
	}

	return nil
}

func WriteKey() {

}

func DeleteKey() {

}

func CheckAndCreateStorageFile() error {
	storageFileExists, err := checkStorageFileExists()

	if err != nil {
		return err
	}

	if !storageFileExists {
		err = createStorageFile()

		if err != nil {
			return err
		}
	}

	return nil
}
