package balancer

import (
	"errors"
	"math/rand"
)




type WeightBalance struct {

}
//TODO
//权重算法还未实现的，试了下接口的使用

func init() {
	RegisterBalance("weight",&WeightBalance{})
}

func (b *WeightBalance) DoBalance(insts []*Instance) (err error, inst *Instance) {
	if len(insts) == 0 {
		err = errors.New("instance do not exist")
		return
	}
	inst = insts[rand.Intn(len(insts))]
	return
}
