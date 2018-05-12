// Resources
//
// A Tour of Go - https://tour.golang.org/
// Inheritance in Go - https://hackthology.com/object-oriented-inheritance-in-go.html
// crypto/rand - https://golang.org/pkg/crypto/rand/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var numCPU = runtime.NumCPU()

// Keeps track of which villains have been fought
var villainMap = SafeVillainMap{fought: make(map[string]bool)}

type SafeVillainMap struct {
	fought map[string]bool
	mux    sync.Mutex
}

func (m *SafeVillainMap) Add(villain string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.fought[villain] = true
}

func (m *SafeVillainMap) Contains(villain string) bool {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.fought[villain]
}

func main() {
	fmt.Println("** Heroes **")
	for _, v := range heroes {
		fmt.Printf("Name: %v\n\tAttack Power: %v\n\tDefense: %v\n", v.name, v.attackPower, v.defense)
	}

	fmt.Println("-- Villains --")
	for _, v := range villains {
		fmt.Printf("Name: %v\n\tAttack Power: %v\n\tDefense: %v\n", v.name, v.attackPower, v.defense)
	}

	// TODO
	// incorporate advantages and disadvantages for each character
	// teaming up against villains (villain "hero capacity")
	// special attacks and randomness for each battle

	// Other Ideas
	// Randomly add general enemies to a slice which the heroes iterate over and fight
}

func SaveTheWolrd() {
	// TODO have the heroes fight the villains concurrently (do attack power damage to defense and vice versa)
}

type Character struct {
	name        string
	attackPower int
	defense     int
}

var heroes = []Character{
	Character{
		name:        "Thor",
		attackPower: 20,
		defense:     50,
	},
	Character{
		name:        "Iron Man",
		attackPower: 15,
		defense:     45,
	},
}

var villains = []Character{
	Character{
		name:        "Thanos",
		attackPower: 25,
		defense:     100,
	},
	Character{
		name:        "Ultron",
		attackPower: 15,
		defense:     80,
	},
}
