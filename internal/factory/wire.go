//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"

	"github.com/mr-panta/wire-experiment/internal/mymodule"
	"github.com/mr-panta/wire-experiment/internal/mypackage"
)

func NewDefaultManager(a mymodule.AInt, b mypackage.BInt) (*Manager, error) {
	wire.Build(
		NewManager,
		mymodule.NewAObject,
		mypackage.NewBObject,
		mypackage.NewCObject,
		wire.Bind(new(mymodule.IAObject), new(*mymodule.AObject)),
		wire.Bind(new(mypackage.IBObject), new(*mypackage.BObject)),
		wire.Bind(new(mypackage.ICObject), new(*mypackage.CObject)),
	)
	return nil, nil
}
