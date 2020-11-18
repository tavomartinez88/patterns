// builder_wrapper is a package wrapper for create users
package builderwrapper

import (
	"github.com/tavomartinez88/go/BuilderUser/builders"
	"github.com/tavomartinez88/go/BuilderUser/domain"
)

type GeneratorUser struct {}

func (bu *GeneratorUser) BuildUser(constructor builders.BuilderUser, data map[string]string) *domain.User {
		return constructor.Build(data)
}










