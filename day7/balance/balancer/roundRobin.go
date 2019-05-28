package balancer

import (
	"errors"
)

type RoundRobin struct {
	CurIndex int
}


func init() {
	RegisterBalance("round",&RoundRobin{})
}

func (b *RoundRobin) DoBalance(insts []*Instance) (err error, inst *Instance) {
	if len(insts) == 0 {
		err = errors.New("instance do not exist")
		return
	}
	lens :=len(insts)
	if b.CurIndex>=lens{
		b.CurIndex=0
	}
	inst=insts[b.CurIndex]
	b.CurIndex++
	return
}
