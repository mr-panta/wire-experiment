package mymodule

type IAObject interface {
	AFunc()
}

type AObject struct {
	a int
}

type AInt int

func NewAObject(a AInt) *AObject {
	return &AObject{
		a: int(a),
	}
}

func (*AObject) AFunc() {}
