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
	numHeroes := len(heroes)
	c := make([]chan string, numHeroes)

	fmt.Println("FIGHT!\n")
	for i, v := range heroes {
		c[i] = make(chan string)
		go SaveTheWolrd(i*len(villains)/numHeroes, (i+1)*len(villains)/numHeroes, v, c[i])
	}

	for i := range c {
		for s := range c[i] {
			fmt.Println(s)
		}
	}
	// TODO
	// incorporate advantages and disadvantages for each character
	// teaming up against villains (villain "hero capacity")
	// special attacks and randomness for each battle

	// Other Ideas
	// Randomly add general enemies to a slice which the heroes iterate over and fight

	fmt.Println("\nFinal Stats")
	PrintStats(heroes, "** Heroes **")
	PrintStats(villains, "-- Villains --")
}

func SaveTheWolrd(i, n int, hero Character, c chan string) {
	defer close(c)
	// TODO have the heroes fight the villains concurrently (do attack power damage to defense and vice versa)
	for ; i < n; i++ {
		villains[i].health -= hero.attackPower
		c <- fmt.Sprintf("%v does %v damage to %v", hero.name, hero.attackPower, villains[i].name)
	}

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
	Character{
		name:        "Thug",
		attackPower: 2,
		defense:     10,
		health:      20,
	},
	Character{
		name:        "Criminal",
		attackPower: 3,
		defense:     8,
		health:      15,
	},
	Character{
		name:        "Goon",
		attackPower: 31,
		defense:     5,
		health:      10,
	},
}
