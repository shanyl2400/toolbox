package service

import (
	"io"
	"path"
	"toolbox/internal/errors"
	"toolbox/internal/storage"
	"toolbox/internal/utils"
)

var (
	profileExtensions = []string{
		".jpg", ".jpeg", ".png", ".gif",
	}
)

type IProfilesService interface {
	Upload(fileName string, reader io.Reader) (string, error)
	Contains(name string) bool
	Download(name string, w io.Writer) error
}

type profilesService struct {
	storage storage.IStorage
}

func (p *profilesService) Upload(fileName string, reader io.Reader) (string, error) {
	ext := path.Ext(fileName)
	if ext == "" {
		return "", errors.ErrUnknownAssetsType
	}
	if !p.checkExtension(ext) {
		return "", errors.ErrUnsupportedAssetsType
	}
	assetsName := utils.NewID() + ext
	err := p.storage.Save(assetsName, reader)
	if err != nil {
		return "", errors.ErrSaveAssetsFailed
	}

	return assetsName, nil
}

func (p *profilesService) Contains(name string) bool {
	return p.storage.Contains(name)
}

func (p *profilesService) Download(name string, w io.Writer) error {
	err := p.storage.Load(name, w)
	if err != nil {
		return errors.ErrLoadAssetsFailed
	}
	return nil
}

func (p *profilesService) checkExtension(extension string) bool {
	for _, ext := range profileExtensions {
		if extension == ext {
			return true
		}
	}
	return false
}

func NewProfilesService() IProfilesService {
	return &profilesService{
		storage: storage.Get(),
	}
}
