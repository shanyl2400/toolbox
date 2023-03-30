package repository

import (
	"go.uber.org/zap"
	"sync"
	"toolbox/internal/logs"
	"toolbox/internal/repository/boltdb"
	"toolbox/internal/repository/model"
)

type UsersRepository interface {
	Set(u *model.User) error
	Delete(userName string) error

	Get(userName string) (*model.User, error)
	List() ([]*model.User, error)
}

var (
	_usersRepo     UsersRepository
	_usersRepoOnce sync.Once
)

func GetUsersRepository() UsersRepository {
	var err error
	_usersRepoOnce.Do(func() {
		_usersRepo, err = boltdb.NewUsersRepository()
		if err != nil {
			logs.Error("create users repository failed",
				zap.Error(err))
		}
	})
	return _usersRepo
}
