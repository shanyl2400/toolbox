package boltdb

import (
	"go.etcd.io/bbolt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"toolbox/internal/errors"
	"toolbox/internal/repository/model"
	"toolbox/internal/utils"
)

type ColumnsRepository struct {
	client *Client
}

func (b *ColumnsRepository) Set(c *model.Column) error {
	return b.client.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(columnsBucket))
		if c.Name == "" {
			//添加且id为空
			c.Name = bson.NewObjectId().Hex()
			c.CreatedAt = time.Now().Unix()
		} else if b.Get([]byte(c.Name)) == nil {
			//添加
			c.CreatedAt = time.Now().Unix()
		}
		c.UpdatedAt = time.Now().Unix()
		data, err := utils.Encode(c)
		if err != nil {
			return err
		}
		return b.Put([]byte(c.Name), data)
	})
}

func (b *ColumnsRepository) Delete(name string) error {
	return b.client.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(columnsBucket))
		return b.Delete([]byte(name))
	})
}

func (b *ColumnsRepository) Get(name string) (*model.Column, error) {
	column := new(model.Column)
	err := b.client.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(columnsBucket))
		data := b.Get([]byte(name))
		if data == nil {
			return errors.ErrNil
		}
		return utils.Decode(data, column)
	})
	if err != nil {
		return nil, err
	}
	return column, nil
}

func (b *ColumnsRepository) List() ([]*model.Column, error) {
	columns := make([]*model.Column, 0)
	err := b.client.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(columnsBucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			column := new(model.Column)

			err := utils.Decode(v, column)
			if err != nil {
				return err
			}
			columns = append(columns, column)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return columns, nil
}

func NewColumnsRepository() (*ColumnsRepository, error) {
	err := GetClient().db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(columnsBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ColumnsRepository{
		client: GetClient(),
	}, nil
}
