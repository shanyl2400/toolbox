package boltdb

import (
	"go.etcd.io/bbolt"
	"time"
	"toolbox/internal/errors"
	"toolbox/internal/repository/model"
	"toolbox/internal/utils"
)

const (
	usersBucket   = "users"
	toolsBucket   = "tools"
	columnsBucket = "columns"
)

type UsersRepository struct {
	client *Client
}

func (b *UsersRepository) Set(u *model.User) error {
	return b.client.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(usersBucket))
		if b.Get([]byte(u.UserName)) == nil {
			//添加
			u.CreatedAt = time.Now().Unix()
		}
		u.UpdatedAt = time.Now().Unix()
		data, err := utils.Encode(u)
		if err != nil {
			return err
		}
		return b.Put([]byte(u.UserName), data)
	})
}

func (b *UsersRepository) Delete(userName string) error {
	return b.client.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(usersBucket))
		return b.Delete([]byte(userName))
	})
}

func (b *UsersRepository) Get(userName string) (*model.User, error) {
	user := new(model.User)
	err := b.client.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(usersBucket))
		data := b.Get([]byte(userName))
		if data == nil {
			return errors.ErrNil
		}
		return utils.Decode(data, user)
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *UsersRepository) List() ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := b.client.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(usersBucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			user := new(model.User)
			err := utils.Decode(v, user)
			if err != nil {
				return err
			}
			users = append(users, user)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUsersRepository() (*UsersRepository, error) {
	err := GetClient().db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(usersBucket))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &UsersRepository{
		client: GetClient(),
	}, nil
}
