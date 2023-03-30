package storage

import (
	errors2 "errors"
	"go.uber.org/zap"
	"io"
	"os"
	"toolbox/internal/config"
	"toolbox/internal/errors"
	"toolbox/internal/logs"
)

type IStorage interface {
	Save(name string, r io.Reader) error
	Load(name string, w io.Writer) error
	Contains(name string) bool
}

type LocalStorage struct {
	path string
}

func (l *LocalStorage) Save(name string, r io.Reader) error {
	path := l.path + "/" + name
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logs.Warn("open file failed",
			zap.Error(err),
			zap.String("path", path))
		return err
	}
	_, err = io.Copy(f, r)
	if err != nil {
		logs.Warn("copy data failed",
			zap.Error(err),
			zap.String("path", path))
		return err
	}
	return nil
}

func (l *LocalStorage) Load(name string, w io.Writer) error {
	if !l.Contains(name) {
		return errors.ErrAssetsNotExists
	}
	path := l.path + "/" + name
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		logs.Warn("open file failed",
			zap.Error(err),
			zap.String("path", path))
		return err
	}
	_, err = io.Copy(w, f)
	if err != nil {
		logs.Warn("copy data failed",
			zap.Error(err),
			zap.String("path", path))
		return err
	}
	return nil
}

func (l *LocalStorage) Contains(name string) bool {
	path := l.path + "/" + name
	if _, err := os.Stat(path); errors2.Is(err, os.ErrNotExist) {
		// path does not exist
		return false
	}
	return true
}

func Get() IStorage {
	return &LocalStorage{
		path: config.Get().AssetsPath,
	}
}
