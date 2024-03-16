package issue46893

import "runtime/debug"

type Panicker interface {
	Panic([]byte) ([]byte, error)
}

type PanickerImpl struct {
}

func (p *PanickerImpl) Panic(b []byte) ([]byte, error) {
	return make([]byte, len(b)), nil
}

func Init() {
	debug.SetTraceback("crash")
}
