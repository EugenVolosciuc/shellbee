package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func saveStorageFile(data []byte) error {
	filePath, err := getStorageFileLocation()

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, []byte(data), 0644)

	if err != nil {
		return err
	}

	return nil
}

func createStorageFile() error {
	filePath, err := getStorageFileLocation()

	if err != nil {
		return err
	}

	jsonStr, _ := json.Marshal(StorageFileData{make(AliasesMap)})

	err = os.MkdirAll(filepath.Dir(filePath), 0755)

	if err != nil {
		return err
	}

	err = saveStorageFile(jsonStr)

	if err != nil {
		return err
	}

	return nil
}

func WriteKey(key, command string) error {
	storageFileData, err := readStorageFile()

	if err != nil {
		return err
	}

	storageFileData.Aliases[key] = command

	jsonStr, _ := json.Marshal(storageFileData)

	saveStorageFile(jsonStr)

	return nil
}

func DeleteKey(key string) error {
	storageFileData, err := readStorageFile()

	if err != nil {
		return err
	}

	_, ok := storageFileData.Aliases[key]

	if !ok {
		return &KeyNotFoundError{key}
	}

	delete(storageFileData.Aliases, key)

	jsonStr, _ := json.Marshal(storageFileData)

	saveStorageFile(jsonStr)

	return nil
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
