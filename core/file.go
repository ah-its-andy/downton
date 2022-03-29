package core

import (
	"io"
	"os"
	"sync"
)

type FileProvider interface {
	Reader() (io.Reader, error)
	Writer() (io.Writer, error)
	Close() error
}

type LocalFileProvider struct {
	path string

	file *os.File

	once sync.Once
}

func NewLocalFileProvider(path string) *LocalFileProvider {
	return &LocalFileProvider{
		path: path,
	}
}

func (provider *LocalFileProvider) openFile() error {
	var err error
	provider.once.Do(func() {
		file, innerErr := os.Open(provider.path)
		if innerErr == nil {
			provider.file = file
		} else {
			err = innerErr
		}
	})
	if err != nil {
		return err
	}
	return nil
}

func (provider *LocalFileProvider) Reader() (io.Reader, error) {
	err := provider.openFile()
	if err != nil {
		return nil, err
	}
	return provider.file, nil
}
func (provider *LocalFileProvider) Writer() (io.Writer, error) {
	err := provider.openFile()
	if err != nil {
		return nil, err
	}
	return provider.file, nil
}

func (provider *LocalFileProvider) Close() error {
	if provider.file != nil {
		return provider.file.Close()
	}
	return nil
}
