package factory

import (
	"fmt"
	"reflect"

	"github.com/mr-panta/wire-experiment/internal/mymodule"
	"github.com/mr-panta/wire-experiment/internal/mypackage"
)

type Manager struct {
	aObj mymodule.IAObject
	bObj mypackage.IBObject
}

func NewManager(aObj mymodule.IAObject, bObj mypackage.IBObject) *Manager {
	return &Manager{
		aObj: aObj,
		bObj: bObj,
	}
}

func (m *Manager) String() string {
	return fmt.Sprintf("a:%+v, b:%+v", reflect.TypeOf(m.aObj), reflect.TypeOf(m.bObj))
}
