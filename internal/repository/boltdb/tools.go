package boltdb

import (
	"go.etcd.io/bbolt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"toolbox/internal/errors"
	"toolbox/internal/repository/model"
	"toolbox/internal/utils"
)

type ToolsRepository struct {
	client *Client
}

func (b *ToolsRepository) Set(t *model.Tool) error {
	return b.client.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(toolsBucket))
		if t.ID == "" {
			//添加且id为空
			t.ID = bson.NewObjectId().Hex()
			t.CreatedAt = time.Now().Unix()
		} else if b.Get([]byte(t.ID)) == nil {
			//添加
			t.CreatedAt = time.Now().Unix()
		}
		t.UpdatedAt = time.Now().Unix()
		data, err := utils.Encode(t)
		if err != nil {
			return err
		}
		return b.Put([]byte(t.ID), data)
	})
}

func (b *ToolsRepository) Delete(id string) error {
	return b.client.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(toolsBucket))
		return b.Delete([]byte(id))
	})
}

func (b *ToolsRepository) Get(id string) (*model.Tool, error) {
	tool := new(model.Tool)
	err := b.client.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(toolsBucket))
		data := b.Get([]byte(id))
		if data == nil {
			return errors.ErrNil
		}
		return utils.Decode(data, tool)
	})

	if err != nil {
		return nil, err
	}
	return tool, nil
}

func (b *ToolsRepository) List() ([]*model.Tool, error) {
	tools := make([]*model.Tool, 0)
	err := b.client.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(toolsBucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tool := new(model.Tool)
			err := utils.Decode(v, tool)
			if err != nil {
				return err
			}
			tools = append(tools, tool)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tools, nil
}

func NewToolsRepository() (*ToolsRepository, error) {
	err := GetClient().db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(toolsBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ToolsRepository{
		client: GetClient(),
	}, nil
}
