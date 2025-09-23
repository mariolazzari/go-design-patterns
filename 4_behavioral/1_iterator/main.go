package main

import (
	"fmt"
)

// Iterator Interface
type Iterator interface {
	hasNext() bool
	next() interface{}
}

// Collection Interface
type Collection interface {
	getIterator() Iterator
}

type IceCreamFlavor struct {
	name string
}

// Concrete Iterator
type IceCreamIterator struct {
	flavors []IceCreamFlavor
	current int
}

func (i *IceCreamIterator) hasNext() bool {
	return i.current < len(i.flavors)
}

func (i *IceCreamIterator) next() interface{} {

	flavor := i.flavors[i.current]
	i.current++
	return flavor
}

// Concrete Collection
type IceCreamShop struct {
	flavors []IceCreamFlavor
}

func (s *IceCreamShop) getIterator() Iterator {
	return &IceCreamIterator{
		flavors: s.flavors,
	}
}

func main() {
	shop := &IceCreamShop{
		flavors: []IceCreamFlavor{
			{"Chocolate"},
			{"Vanilla"},
			{"Pistachio"},
			{"Cookies & Cream"},
		},
	}

	iterator := shop.getIterator()
	for iterator.hasNext() {
		flavor := iterator.next().(IceCreamFlavor)
		fmt.Printf("%s\n", flavor.name)
	}
}
