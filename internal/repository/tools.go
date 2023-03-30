package repository

import (
	"go.uber.org/zap"
	"sync"
	"toolbox/internal/logs"
	"toolbox/internal/repository/boltdb"
	"toolbox/internal/repository/model"
)

type ToolsRepository interface {
	Set(t *model.Tool) error

	Delete(id string) error
	Get(id string) (*model.Tool, error)
	List() ([]*model.Tool, error)
}

var (
	_toolsRepo     ToolsRepository
	_toolsRepoOnce sync.Once
)

func GetToolsRepository() ToolsRepository {
	var err error
	_toolsRepoOnce.Do(func() {
		_toolsRepo, err = boltdb.NewToolsRepository()
		if err != nil {
			logs.Error("create tools repository failed",
				zap.Error(err))
		}
	})
	return _toolsRepo
}
