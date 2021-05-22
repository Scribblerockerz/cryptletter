package attachment

import (
	"errors"
	"github.com/Scribblerockerz/cryptletter/pkg/utils"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type localTempHandler struct {
	list        utils.DismissiveList
	defaultTTL  int64
	storagePath string
}

func (l localTempHandler) getStorageDir() (string, error) {
	storagePath := l.storagePath
	if !path.IsAbs(storagePath) {
		storagePath, _ = filepath.Abs(path.Join(utils.GetConfigRootDir(), storagePath))
	}

	stat, err := os.Stat(storagePath)
	if os.IsNotExist(err) {
		err = os.Mkdir(l.storagePath, 0755)
		return "", err
	}

	if !stat.IsDir() {
		return "", errors.New("storage directory can not be a file")
	}

	return storagePath, nil
}

//Put will place data into the storage directory
func (l localTempHandler) Put(fileData string) (string, error) {
	storageDir, err := l.getStorageDir()
	if err != nil {
		return "", err
	}

	tmpFile, err := ioutil.TempFile(storageDir, "att-")
	if err != nil {
		return "", err
	}

	defer tmpFile.Close()

	if _, err = tmpFile.Write([]byte(fileData)); err != nil {
		return "", errors.New("failed to write to temporary file")
	}

	identifier := path.Base(tmpFile.Name())

	// Add file to the list of tracked resources
	err = l.list.Add(identifier, l.defaultTTL)
	if err != nil {
		return "", err
	}

	return identifier, nil
}

//Get will retrieve file data by a given identifier
func (l localTempHandler) Get(identifier string) (string, error) {
	// Add file to the list of tracked resources
	if !l.list.Has(identifier) {
		return "", errors.New("unable to read file by identifier")
	}

	storageDir, err := l.getStorageDir()
	if err != nil {
		return "", err
	}

	filePath := path.Join(storageDir, identifier)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.New("unable to read file by identifier")
	}

	return string(data), nil
}

//Delete will remove stored files by identifier
func (l localTempHandler) Delete(identifier string) error {
	// Delete identifier from tracked resources
	l.list.Del(identifier)

	storageDir, err := l.getStorageDir()
	if err != nil {
		return err
	}

	filePath := path.Join(storageDir, identifier)

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	}

	err = os.Remove(filePath)
	if err != nil {
		return errors.New("unable to remove file by identifier")
	}

	return nil
}

//SetTTL will update the TTL of an identifier
func (l localTempHandler) SetTTL(identifier string, ttl int64) error {
	return l.list.Set(identifier, ttl)
}

//Cleanup will sync the known file identifiers with the unknown and clean them up
func (l localTempHandler) Cleanup() error {
	// Get storage directory
	// Get all known files from the list.All()
	// Iterate over the storage directory and remove everything which is not on the list.All()

	trackedIdentifier, err := l.list.All()
	if err != nil {
		return err
	}

	storageDir, err := l.getStorageDir()
	if err != nil {
		return err
	}

	err = filepath.WalkDir(storageDir, func(p string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if utils.ListContainsString(trackedIdentifier, info.Name()) {
			return nil
		}

		err = os.Remove(path.Join(storageDir, info.Name()))
		if err != nil {
			return errors.New("unable to remove timed out file " + info.Name())
		}

		return nil
	})

	return nil
}

//DropAll will clear the storage dir and the list
func (l localTempHandler) DropAll() error {
	storageDir, err := l.getStorageDir()
	if err != nil {
		return err
	}

	_, err = os.Stat(storageDir)
	if !os.IsNotExist(err) {
		err = os.Remove(storageDir)
		if err != nil {
			return err
		}
	}

	return l.list.Drp()
}

func NewLocalTempHandler(defaultTTL int64, storagePath string) Handler {
	return &localTempHandler{
		list:        utils.NewDismissiveList("local-files"),
		defaultTTL:  defaultTTL,
		storagePath: storagePath,
	}
}
