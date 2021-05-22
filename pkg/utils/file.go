package utils

import (
	"bytes"
	"io/fs"
	"time"
)

type InMemoryFile struct {
	fs.File
	InMemoryFileInfo InMemoryFileInfo // Size() will return an inaccurate value, since we modified it
	Buf *bytes.Buffer
}

func (im InMemoryFile) Stat() (fs.FileInfo, error) {
	return im.InMemoryFileInfo, nil
}

func (im InMemoryFile) Read(b []byte) (int, error) {
	return im.Buf.Read(b)
}

func (im InMemoryFile) Close() error {
	return nil
}

type InMemoryFileInfo struct {
	fs.FileInfo
	FileInfoRef fs.FileInfo
	FileSize int64
}

func (fi InMemoryFileInfo) Name() string {
	return fi.FileInfoRef.Name()
}

func (fi InMemoryFileInfo) Size() int64 {
	return fi.FileSize
}

func (fi InMemoryFileInfo) Mode() fs.FileMode {
	return fi.FileInfoRef.Mode()
}

func (fi InMemoryFileInfo) ModTime() time.Time {
	return fi.FileInfoRef.ModTime()
}

func (fi InMemoryFileInfo) IsDir() bool {
	return fi.FileInfoRef.IsDir()
}
