package factory

type AObject interface{}
type BObject interface{}

type Manager struct {
	aObj AObject
	bObj BObject
}

func NewManager(aObj AObject, bObj BObject) *Manager {
	return &Manager{
		aObj: aObj,
		bObj: bObj,
	}
}
