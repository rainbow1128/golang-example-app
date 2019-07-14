// +build wireinject

package oauth

import (
	"github.com/aristat/golang-gin-oauth2-example-app/app/provider"
	"github.com/aristat/golang-gin-oauth2-example-app/app/session"
	"github.com/google/wire"
)

// Build
func Build() (*OAuth, func(), error) {
	panic(wire.Build(ProviderProductionSet, session.ProviderProductionSet, provider.AwareProductionSet))
}
