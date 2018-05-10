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
	fmt.Println("Heroes")
	for _, v := range heroes {
		fmt.Printf("Name: %v, Attack Power: %v\n", v.name, v.attackPower)
	}

	fmt.Println("Villains")
	for _, v := range villains {
		fmt.Printf("Name: %v, Defense: %v\n", v.name, v.defense)
	}

	// TODO
	// have the heroes fight the villains concurrently (do attack power damage to defense)
	// later incorporate advantages and disadvantages and teaming up against villains (villain "hero capacity")
	// combine character attributes into one struct for reuse
}

type avenger struct {
	name        string
	attackPower int
}

var heroes = []avenger{
	avenger{
		name:        "Thor",
		attackPower: 20,
	},
	avenger{
		name:        "Iron Man",
		attackPower: 15,
	},
}

type villain struct {
	name    string
	defense int
}

var villains = []villain{
	villain{
		name:    "Thanos",
		defense: 100,
	},
	villain{
		name:    "Ultron",
		defense: 80,
	},
}
