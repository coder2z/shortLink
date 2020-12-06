package server

import (
	"context"
	_map "shortLink/internal/master/map"
	"shortLink/internal/master/model"
	"shortLink/internal/master/model/constant"
	string_plus "shortLink/pkg/string"
	"time"
)

func GetUrlByKey(id _map.IdMap) (string, error) {
	if res := model.Redis.Get().Get(context.Background(), id.Id); res.Val() != constant.StrNil {
		return res.Val(), nil
	}
	urls := new(model.Urls)
	if err := urls.GetByWhere(map[string]interface{}{"url_key=?": id.Id}); err != nil {
		return constant.StrNil, err
	}
	if res := model.Redis.Get().Set(context.Background(), urls.UrlKey, urls.ToUrl, time.Hour); res.Err() != nil {
		return urls.ToUrl, res.Err()
	}
	return urls.ToUrl, nil
}

func AddUrl(add _map.AddLink) (string, error) {
	urls := model.Urls{
		UrlKey:   string_plus.Get(),
		ToUrl: add.Url,
	}
	return urls.UrlKey, urls.Add()
}
