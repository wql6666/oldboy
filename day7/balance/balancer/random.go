package balancer

import (
	"errors"
	"math/rand"
)

type Random struct {

}

func init() {
	RegisterBalance("random",&Random{})
}

func (b *Random) DoBalance(insts []*Instance) (err error, inst *Instance) {
	if len(insts) == 0 {
		err = errors.New("instance do not exist")
		return
	}
	inst = insts[rand.Intn(len(insts))]
	return
}
