/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/12/6 15:47
 */
package model

import "github.com/jinzhu/gorm"

type Urls struct {
	gorm.Model
	UrlKey string `gorm:"url_key;unique;not null"`
	ToUrl  string `gorm:"to_url;unique;not null"`
}

func (Urls) TableName() string {
	return "urls"
}

func (a *Urls) Add() error {
	return MainDB.Table(a.TableName()).Create(a).Error
}

func (a *Urls) CreateOrUpdate() {

}

func (a *Urls) Del(wheres map[string]interface{}) error {
	db := MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	return db.Delete(a).Error
}

func (a *Urls) GetAll(data *[]Urls) (err error) {
	err = MainDB.Table(a.TableName()).Find(&data).Error
	return
}

func (a *Urls) Get(start int64, size int64, data *[]Urls, wheres map[string]interface{}) (total int64, err error) {
	db := MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	err = db.Limit(size).Offset(start).Find(&data).Error
	err = db.Count(&total).Error
	return
}

func (a *Urls) GetById() error {
	return a.GetByWhere(map[string]interface{}{
		"id=?": a.ID,
	})
}

func (a *Urls) UpdateById() error {
	return MainDB.Table(a.TableName()).Where("id=?", a.ID).Update(a).Error
}

func (a *Urls) GetByWhere(wheres map[string]interface{}) error {
	db := MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	return db.Find(a).Error
}
