// Resources
//
// A Tour of Go - https://tour.golang.org/
// Inheritance in Go - https://hackthology.com/object-oriented-inheritance-in-go.html
// crypto/rand - https://golang.org/pkg/crypto/rand/

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

func PrintStats(c CharacterList, header string) {
	fmt.Println(header)
	for _, v := range c.list {
		fmt.Printf("Name: %v\n\tAttack Power: %v\n\tDefense: %v\n\tHealth: %v\n", v.name, v.attackPower, v.defense, v.health)
	}
}

func main() {
	PrintStats(heroes, "** Heroes **")
	PrintStats(villains, "-- Villains --")
	numHeroes := len(heroes.list)
	numVillains := len(villains.list)
	c := make([]chan string, numHeroes)

	fmt.Println("FIGHT!\n")
	for i, _ := range heroes.list {
		c[i] = make(chan string)
		go SaveTheWorld(i*numVillains/numHeroes, (i+1)*numVillains/numHeroes, heroes.list[i], c[i])
	}

	for i := range c {
		for s := range c[i] {
			fmt.Println(s)
		}
	}
	// TODO Ideas
	// incorporate advantages and disadvantages for each character
	// teaming up against villains (villain "hero capacity")
	// special attacks
	// Randomly add general enemies to a slice which the heroes iterate over and fight
	// while at least one hero is alive and enemies remain
	// print living heroes and villains

	fmt.Println("\nFinal Stats")
	PrintStats(heroes, "** Heroes **")
	PrintStats(villains, "-- Villains --")
}

func SaveTheWorld(i, n int, hero *Character, c chan string) {
	defer close(c)

	for ; i < n; i++ {
		if villains.list[i].health > 0 && hero.health > 0 {
			Battle(hero, villains.list[i], c)
		}

		if hero.health > 0 && villains.list[i].health > 0 {
			Battle(villains.list[i], hero, c)
		}
	}
}

func Battle(attacker, defender *Character, c chan string) {
	damage := CalculateDamage(attacker.attackPower)
	defender.health -= damage
	c <- fmt.Sprintf("%v does %v damage to %v", attacker.name, damage, defender.name)
}

// Adapted from: https://stackoverflow.com/a/32350135/8094831
func CalculateDamage(i int) (damage int) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	//fmt.Printf("Here is a random %T in [0,%v) : %d\n", n, n, i)
	return int(n)
}

type Character struct {
	name        string
	attackPower int
	defense     int
	health      int
}

type CharacterList struct {
	list []*Character
}

var heroes = CharacterList{
	[]*Character{
		&Character{
			name:        "Thor",
			attackPower: 20,
			defense:     50,
			health:      70,
		},
		&Character{
			name:        "Iron Man",
			attackPower: 15,
			defense:     45,
			health:      60,
		},
		&Character{
			name:        "Spider-Man",
			attackPower: 10,
			defense:     30,
			health:      50,
		},
	},
}

var villains = CharacterList{
	[]*Character{
		&Character{
			name:        "Thanos",
			attackPower: 25,
			defense:     100,
			health:      200,
		},
		&Character{
			name:        "Ultron",
			attackPower: 15,
			defense:     80,
			health:      150,
		},
		&Character{
			name:        "Thug",
			attackPower: 2,
			defense:     10,
			health:      20,
		},
		&Character{
			name:        "Criminal",
			attackPower: 3,
			defense:     8,
			health:      15,
		},
		&Character{
			name:        "Goon",
			attackPower: 1,
			defense:     5,
			health:      10,
		},
	},
}
