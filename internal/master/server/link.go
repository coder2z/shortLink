package server

import (
	"context"
	_map "shortLink/internal/master/map"
	"shortLink/internal/master/model"
)

func GetUrlByKey(id _map.IdMap) (string, error) {
	res := model.Redis.Get().Get(context.Background(), id.Id)
	if res.Err() != nil {
		return "", res.Err()
	}
	return res.Val(), nil
}

func AddUrl(add _map.AddLink) error {
	return model.Redis.Get().Set(context.Background(), "test", add.Url, 0).Err()
}