// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package http

import (
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/db"
	"github.com/aristat/golang-example-app/app/db/repo"
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/graphql"
	"github.com/aristat/golang-example-app/app/grpc"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/app/oauth"
	"github.com/aristat/golang-example-app/app/resolver"
	"github.com/aristat/golang-example-app/app/session"
	"github.com/aristat/golang-example-app/app/tracing"
	"github.com/aristat/golang-example-app/app/users"
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
	sessionConfig, cleanup3, err := session.Cfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup4, err := session.Provider(context, sessionConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup5, err := logger.ProviderCfg(viper)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup6, err := logger.Provider(context, loggerConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbConfig, cleanup7, err := db.Cfg(viper)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gormDB, cleanup8, err := db.ProviderGORM(context, zap, dbConfig)
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
	dbManager, cleanup9, err := db.Provider(context, zap, dbConfig, gormDB)
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
	oauthManager, cleanup13, err := oauth.Provider(context, zap, tokenStore, manager, clientStore)
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
	usersRepo, cleanup14, err := repo.NewUsersRepo(gormDB)
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
	repoRepo, cleanup15, err := repo.Provider(usersRepo)
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
	configuration, cleanup16, err := tracing.ProviderCfg(viper)
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
	tracer, cleanup17, err := tracing.Provider(context, configuration, zap)
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
	grpcConfig, cleanup18, err := grpc.Cfg(viper)
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
	poolManager, cleanup19, err := grpc.Provider(context, tracer, zap, grpcConfig)
	if err != nil {
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
		return nil, nil, err
	}
	serviceManagers := users.ServiceManagers{
		Session:     manager,
		DB:          dbManager,
		Oauth:       oauthManager,
		Repo:        repoRepo,
		PoolManager: poolManager,
	}
	usersManager, cleanup20, err := users.Provider(context, zap, serviceManagers)
	if err != nil {
		cleanup19()
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
		return nil, nil, err
	}
	resolverConfig, cleanup21, err := resolver.Cfg(viper)
	if err != nil {
		cleanup20()
		cleanup19()
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
		return nil, nil, err
	}
	managers := resolver.Managers{
		Repo:        repoRepo,
		PollManager: poolManager,
	}
	graphqlConfig, cleanup22, err := resolver.Provider(context, zap, resolverConfig, managers)
	if err != nil {
		cleanup21()
		cleanup20()
		cleanup19()
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
		return nil, nil, err
	}
	config2, cleanup23, err := graphql.Cfg(viper)
	if err != nil {
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
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
		return nil, nil, err
	}
	graphQL, cleanup24, err := graphql.Provider(context, graphqlConfig, zap, config2)
	if err != nil {
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
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
		return nil, nil, err
	}
	httpManagers := Managers{
		session: manager,
		users:   usersManager,
		oauth:   oauthManager,
		graphql: graphQL,
	}
	chiMux, cleanup25, err := Mux(httpManagers, zap, tracer)
	if err != nil {
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
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
		return nil, nil, err
	}
	httpConfig, cleanup26, err := Cfg(viper)
	if err != nil {
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
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
		return nil, nil, err
	}
	http, cleanup27, err := Provider(context, chiMux, zap, httpConfig, httpManagers)
	if err != nil {
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
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
		return nil, nil, err
	}
	return http, func() {
		cleanup27()
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
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
