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

func PrintStats(c []Character, header string) {
	fmt.Println(header)
	for _, v := range c {
		fmt.Printf("Name: %v\n\tAttack Power: %v\n\tDefense: %v\n\tHealth: %v\n", v.name, v.attackPower, v.defense, v.health)
	}
}

func main() {
	PrintStats(heroes, "** Heroes **")
	PrintStats(villains, "-- Villains --")
	// TODO
	// incorporate advantages and disadvantages for each character
	// teaming up against villains (villain "hero capacity")
	// special attacks and randomness for each battle

	// Other Ideas
	// Randomly add general enemies to a slice which the heroes iterate over and fight
}

func SaveTheWolrd() {
	// TODO have the heroes fight the villains concurrently (do attack power damage to defense and vice versa)

	// while at least one hero is alive and enemies remain
	// go fight(hero, enemy slice) --hero iterates over slice doing damage and taking damage
	// print living heroes and villains
}

type Character struct {
	name        string
	attackPower int
	defense     int
	health      int
}

var heroes = []Character{
	Character{
		name:        "Thor",
		attackPower: 20,
		defense:     50,
		health:      70,
	},
	Character{
		name:        "Iron Man",
		attackPower: 15,
		defense:     45,
		health:      60,
	},
}

var villains = []Character{
	Character{
		name:        "Thanos",
		attackPower: 25,
		defense:     100,
		health:      200,
	},
	Character{
		name:        "Ultron",
		attackPower: 15,
		defense:     80,
		health:      150,
	},
}
