// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package http

import (
	"github.com/aristat/golang-oauth2-example-app/app/config"
	"github.com/aristat/golang-oauth2-example-app/app/db"
	"github.com/aristat/golang-oauth2-example-app/app/db/repo"
	"github.com/aristat/golang-oauth2-example-app/app/entrypoint"
	"github.com/aristat/golang-oauth2-example-app/app/logger"
	"github.com/aristat/golang-oauth2-example-app/app/oauth"
	"github.com/aristat/golang-oauth2-example-app/app/session"
	"github.com/aristat/golang-oauth2-example-app/app/users"
)

// Injectors from injector.go:

func Build() (*Http, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup3, err := logger.ProviderCfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup4, err := logger.Provider(context, loggerConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbConfig, cleanup5, err := db.Cfg(viper)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gormDB, cleanup6, err := db.ProviderGORM(context, zap, dbConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup7, err := db.Provider(context, zap, dbConfig, gormDB)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sessionConfig, cleanup8, err := session.Cfg(viper)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sessionManager, cleanup9, err := session.Provider(context, sessionConfig)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauthConfig, cleanup10, err := oauth.Cfg(viper)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tokenStore, cleanup11, err := oauth.TokenStore(oauthConfig)
	if err != nil {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientStore, cleanup12, err := oauth.ClientStore()
	if err != nil {
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauthManager, cleanup13, err := oauth.Provider(context, zap, tokenStore, sessionManager, clientStore)
	if err != nil {
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	managers := users.Managers{
		session: sessionManager,
		db:      manager,
		oauth:   oauthManager,
	}
	usersRepo, cleanup14, err := repo.NewAuthorsRepo(gormDB)
	if err != nil {
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repo2 := &users.Repo{
		Users: usersRepo,
	}
	usersManager, cleanup15, err := users.Provider(context, zap, managers, repo2)
	if err != nil {
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpManagers := Managers{
		session: sessionManager,
		users:   usersManager,
		oauth:   oauthManager,
	}
	chiMux, cleanup16, err := Mux(manager, httpManagers, zap)
	if err != nil {
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpConfig, cleanup17, err := Cfg(viper)
	if err != nil {
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	http, cleanup18, err := Provider(context, chiMux, zap, httpConfig, httpManagers)
	if err != nil {
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return http, func() {
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
