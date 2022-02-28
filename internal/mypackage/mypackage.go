package mypackage

type IBObject interface {
	BFunc()
}

type BObject struct {
	b    int
	cObj ICObject
}

type BInt int

func NewBObject(b BInt, cObj ICObject) *BObject {
	return &BObject{
		b:    int(b),
		cObj: cObj,
	}
}

func (*BObject) BFunc() {}

type ICObject interface{}

type CObject struct{}

func NewCObject() (*CObject, error) {
	return &CObject{}, nil
}
