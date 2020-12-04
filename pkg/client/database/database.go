package database

import (
	"github.com/jinzhu/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewDatabaseClient(o *Options, stopCh <-chan struct{}) (c *Client, err error) {
	db, err := gorm.Open(o.Type, o.GetDSN())

	if err != nil {
		return nil, err
	}
	go func() {
		<-stopCh
		_ = db.Close()
	}()

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTablePrefix + defaultTableName
	}

	if o.Debug {
		db = db.Debug()
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(o.MaxIdleConnections)
	db.DB().SetMaxOpenConns(o.MaxOpenConnections)
	return &Client{db: db}, nil

}

func (c *Client) DB() *gorm.DB {
	if c == nil {
		return nil
	}
	return c.db
}
