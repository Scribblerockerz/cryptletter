package handler

import (
	"errors"
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/attachment"
	"io/ioutil"
	"os"
)

type localTempHandler struct {
}

func (l localTempHandler) Put(fileData string) (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "cryptletter-")
	if err != nil {
		return "", errors.New("unable to create temporary file")
	}

	defer tmpFile.Close()

	fmt.Println("Created File: " + tmpFile.Name())

	if _, err = tmpFile.Write([]byte(fileData)); err != nil {
		return "", errors.New("failed to write to temporary file")
	}

	return tmpFile.Name(), nil
}

func (l localTempHandler) Get(identifier string) (string, error) {
	data, err := ioutil.ReadFile(identifier)
	if err != nil {
		return "", errors.New("unable to read file by identifier")
	}

	return string(data), nil
}

func (l localTempHandler) Delete(identifier string) error {
	_, err := os.Stat(identifier)
	if os.IsNotExist(err) {
		return nil
	}

	err = os.Remove(identifier)
	if err != nil {
		return errors.New("unable to remove file by identifier")
	}

	return nil
}

func NewLocalTempHandler() attachment.Handler {
	return &localTempHandler{}
}
