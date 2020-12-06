package master

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shortLink/internal/master/api/v1/registry"
	"shortLink/internal/master/config"
	"shortLink/internal/master/model"
	myValidator "shortLink/internal/master/validator"
	"shortLink/pkg/client/database"
	"shortLink/pkg/client/redis"
	"shortLink/pkg/log"
	string_plus "shortLink/pkg/string"
	"shortLink/pkg/validator"
)

type Server struct {
	Config *config.Cfg
	Server *http.Server
	err    error
}

func (s *Server) PrepareRun(stopCh <-chan struct{}) (err error) {
	s.initCfg()
	s.initLog()
	s.initDB(stopCh)
	s.initRedis(stopCh)
	s.initHttpServer()
	s.initRouter()
	s.initValidator()
	s.initRandStr()
	return s.err
}

func (s *Server) Run(stopCh <-chan struct{}) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-stopCh
		log.Info(fmt.Sprintf("Shutdown server"))
		_ = s.Server.Shutdown(ctx)
	}()
	log.Info(fmt.Sprintf("Start listening on %s", s.Server.Addr))
	err = s.Server.ListenAndServe()
	return nil
}

func (s *Server) initCfg() {
	if s.err != nil {
		return
	}
	s.Config, s.err = config.TryLoadFromDisk()
}

func (s *Server) initRedis(stopCh <-chan struct{}) {
	if s.err != nil {
		return
	}
	model.Redis, s.err = redis.NewRedisClient(s.Config.Redis, stopCh)
}

func (s *Server) initHttpServer() {
	if s.err != nil {
		return
	}
	s.Server = new(http.Server)
	s.Server.Addr = s.Config.Server.Addr
}

func (s *Server) initLog() {
	if s.err != nil {
		return
	}
	s.err = log.NewLog(s.Config.Log)
}

func (s *Server) initRouter() {
	if s.err != nil {
		return
	}
	s.Server.Handler = registry.Router
}

func (s *Server) initValidator() {
	if s.err != nil {
		return
	}
	s.err = validator.Init(s.Config.Server.Locale, myValidator.RegisterValidation)
}

func (s *Server) initDB(stopCh <-chan struct{}) {
	if s.err != nil {
		return
	}
	var c *database.Client
	log.Info(fmt.Sprintf("db init"))
	c, s.err = database.NewDatabaseClient(s.Config.DB, stopCh)
	log.Info(fmt.Sprintf("db init over"))
	model.MainDB = c.DB()
	s.initMigrate()
}

func (s *Server) initMigrate() {
	if s.err != nil {
		return
	}
	if model.MainDB != nil {
		model.MainDB.AutoMigrate(
			new(model.Urls),
		)
	} else {
		s.err = errors.New("db not init")
	}
}

func (s *Server) initRandStr() {
	if s.err != nil {
		return
	}
	string_plus.New()
}
