package balancer

import "fmt"

type Balancer interface {

	DoBalance([]*Instance)(err error, inst *Instance)
}


type Instance struct {
	Host string
	Port int
}

func (p Instance) String()string {
	return fmt.Sprintf("%s:%d",p.Host,p.Port)
}