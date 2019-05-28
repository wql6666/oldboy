package balancer

import "fmt"

type MgrBalance struct {
	AllBalancer map[string]Balancer
}

var Mgr MgrBalance = MgrBalance{
	AllBalancer: make(map[string]Balancer),
}

func (p *MgrBalance) registerBalance(name string, b Balancer) {
	p.AllBalancer[name] = b
}

func RegisterBalance(name string, b Balancer) {
	Mgr.registerBalance(name, b)
}

func UseMgrBalance(name string) (error, Balancer) {
	b, ok := Mgr.AllBalancer[name]
	if !ok {
		err := fmt.Errorf("no balancer\t%s", name)
		fmt.Println(err)
		return err,nil
	}
	return nil,b
}
