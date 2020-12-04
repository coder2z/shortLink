package model

import (
	"github.com/jinzhu/gorm"
	"shortLink/pkg/client/redis"
)

var (
	MainDB *gorm.DB
	Redis  *redis.Client
)
