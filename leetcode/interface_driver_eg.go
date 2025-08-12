package main

import "fmt"

var (
	driverMap = make(map[string]Driver)
)

type Driver interface {
	m1() string
	m2() int
}

type implInterface struct {
}

func (i *implInterface) m1() string {
	return "m1"
}

func (i *implInterface) m2() int {
	return 2
}

func (i *implInterface) m3() int {
	return 3
}

func GetDriver(name string) Driver {
	return driverMap[name]

}
func main() {
	rk := &implInterface{}
	driverMap["one"] = rk
	driver := GetDriver("one")

	rk2 := &implInterface{}

	fmt.Printf("RK=> %v", driver.m2())
	fmt.Printf("\nRK=> %v", rk2.m3())

}
