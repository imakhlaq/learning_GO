package main

import "sync"

var (
	lock1 sync.Mutex
	lock2 sync.Mutex
)

type resource struct {
	name string
	data string
}

func routine1(data resource) {

	lock1.Lock()
	data.name = "dadada"
	routine2(data)
	data.name = "kaaaa"
	lock1.Unlock()

}

func routine2(data resource) {
	lock2.Lock()
	data.name = "sulpher"
	data.data = "huuuuu"
	lock2.Unlock()
}

func main() {
	r := resource{name: "akhlaq", data: "222"}

	go routine1(r)

	//waiting for both
	select {}
}
