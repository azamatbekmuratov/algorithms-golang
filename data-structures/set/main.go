package main

import "fmt"

//Alias type
type set map[string]struct{}

func (s set) add(v string) {
	s[v] = struct{}{}
}

func (s set) remove(v string) {
	delete(s, v)
}

func (s set) has(v string) bool {
	_, ok := s[v]
	return ok
}

func main() {
	var zoo = set{}
	zoo.add("Walrus")
	zoo.add("Parrot")
	zoo.add("Lion")

	// Adding an existing member to the set again
	zoo.add("Lion")

	zoo.remove("Parrot")

	fmt.Println(zoo)

	// Checking the presence of elements in a set
	fmt.Println(zoo.has("Penguin"))
}
