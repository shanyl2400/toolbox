package service

import (
	"go.uber.org/zap"
	"sync"
	"toolbox/internal/errors"
	"toolbox/internal/logs"
	"toolbox/internal/repository"
	"toolbox/internal/repository/model"
	"toolbox/internal/utils"
)

type IUsersService interface {
	Register(name string, password string) error
	Login(name string, password string) (string, error)

	UpdatePassword(name string, oldPassword string, newPassword string) error
	CheckToken(token string) error

	Get(name string) (*model.UserInfo, error)

	List() ([]*model.UserInfo, error)
}

type userService struct {
	sync.Mutex
}

func (u *userService) Register(name string, password string) error {
	u.Lock()
	defer u.Unlock()
	if len(name) < 8 || len(password) < 8 {
		return errors.ErrUserNameOrPasswordTooShort
	}

	_, err := repository.GetUsersRepository().Get(name)
	if err != errors.ErrNil {
		if err != nil {
			logs.Warn("get user failed",
				zap.Error(err),
				zap.String("username", name))
		}
		return errors.ErrUserNameOccupied
	}
	err = repository.GetUsersRepository().Set(&model.User{
		UserName:     name,
		PasswordHash: utils.Hash(password),
	})
	if err != nil {
		logs.Warn("put user failed",
			zap.Error(err),
			zap.String("username", name))
		return errors.ErrPutUserFailed
	}

	return nil
}

func (u *userService) Login(name string, password string) (string, error) {
	user, err := repository.GetUsersRepository().Get(name)
	if err != nil {
		logs.Warn("get user failed",
			zap.Error(err),
			zap.String("username", name))
		return "", errors.ErrGetUserFailed
	}
	if user.PasswordHash != utils.Hash(password) {
		return "", errors.ErrIncorrectPassword
	}
	token, err := utils.GenToken(user.UserName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *userService) UpdatePassword(name string, oldPassword string, newPassword string) error {
	user, err := repository.GetUsersRepository().Get(name)
	if err != nil {
		logs.Warn("get user failed",
			zap.Error(err),
			zap.String("username", name))
		return errors.ErrGetUserFailed
	}
	if user.PasswordHash != utils.Hash(oldPassword) {
		return errors.ErrIncorrectPassword
	}
	err = repository.GetUsersRepository().Set(&model.User{
		UserName:     name,
		PasswordHash: utils.Hash(newPassword),
	})
	if err != nil {
		logs.Warn("put user failed",
			zap.Error(err),
			zap.String("username", name))
		return errors.ErrPutUserFailed
	}
	return nil
}

func (u *userService) CheckToken(token string) error {
	_, err := utils.ParseToken(token)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Get(name string) (*model.UserInfo, error) {
	user, err := repository.GetUsersRepository().Get(name)
	if err != nil {
		return nil, err
	}
	return &model.UserInfo{UserName: user.UserName}, nil
}

func (u *userService) List() ([]*model.UserInfo, error) {
	users, err := repository.GetUsersRepository().List()
	if err != nil {
		return nil, err
	}
	infos := make([]*model.UserInfo, 0, len(users))
	for i := range users {
		infos[i] = &model.UserInfo{UserName: users[i].UserName}
	}
	return infos, nil
}

func NewUsersService() IUsersService {
	return &userService{}
}
