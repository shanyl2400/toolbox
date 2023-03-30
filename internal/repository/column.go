package repository

import (
	"go.uber.org/zap"
	"sync"
	"toolbox/internal/logs"
	"toolbox/internal/repository/boltdb"
	"toolbox/internal/repository/model"
)

type ColumnsRepository interface {
	Set(c *model.Column) error

	Delete(name string) error
	Get(name string) (*model.Column, error)
	List() ([]*model.Column, error)
}

var (
	_columnsRepo     ColumnsRepository
	_columnsRepoOnce sync.Once
)

func GetColumnsRepository() ColumnsRepository {
	var err error
	_columnsRepoOnce.Do(func() {
		_columnsRepo, err = boltdb.NewColumnsRepository()
		if err != nil {
			logs.Error("create column repository failed",
				zap.Error(err))
		}
	})
	return _columnsRepo
}
