package boltdb

import (
	"go.uber.org/zap"
	"sync"
	"toolbox/internal/config"
	"toolbox/internal/logs"

	"go.etcd.io/bbolt"
)

type Client struct {
	db   *bbolt.DB
	path string
}

func (c *Client) Open() error {
	db, err := bbolt.Open(c.path, 0666, nil)
	if err != nil {
		logs.Error("open boltdb failed",
			zap.Error(err),
			zap.String("path", c.path))
		return err
	}
	c.db = db
	return nil
}

func (c *Client) Close() {
	if c.db != nil {
		c.db.Close()
	}
}

var (
	_client     *Client
	_clientOnce sync.Once
)

func GetClient() *Client {
	_clientOnce.Do(func() {
		_client = &Client{
			path: config.Get().BoltDBPath,
		}
	})

	return _client
}
