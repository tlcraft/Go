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
var numHeroes = len(heroes.list)
var numVillains = len(villains.list)
var numBosses = len(majorVillains.list)

// Keeps track of which villains have been fought
type SafeEngagedList struct {
	engagedFighters []*Character
	capacity        int
	mux             sync.Mutex
}

func (m *SafeEngagedList) Add(fighter *Character) {
	m.mux.Lock()
	defer m.mux.Unlock()
	if len(m.engagedFighters) < m.capacity {
		m.engagedFighters = append(m.engagedFighters, fighter)
	}
}

// TODO Remove fighter from list

func PrintStats(c CharacterList, header string) {
	fmt.Println(header)
	for _, v := range c.list {
		fmt.Printf("Name: %v\n\tAttack Power: %v\n\tDefense: %v\n\tHealth: %v\n", v.name, v.attackPower, v.defense, v.health)
	}
}

func HeroesVsVillains() {
	PrintStats(heroes, "** Heroes **")
	PrintStats(villains, "-- Villains --")

	c := make([]chan string, numHeroes)

	fmt.Println("FIGHT!\n")

	for i, _ := range heroes.list {
		c[i] = make(chan string)
		go SaveTheWorld(i*numVillains/numHeroes, (i+1)*numVillains/numHeroes, heroes.list[i], villains.list, c[i])
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

func BossFight() {
	fmt.Println("\nTesting out a Boss Character idea")
	// TODO New Direction
	// Create a list of BossCharacters
	// Start Go routines to send heroes to fight the bosses
	// Engage heroes with boss characters if the capacity is not full
	// Fight until death and then continue iterating over boss array until one side has won
	// When heroes free up send them on to fight again after they recover some amount of health

	ch := make(chan string)

	fmt.Println(testBoss.character.name)
	fmt.Println(len(testBoss.fighterList.engagedFighters))
	testBoss.fighterList.Add(&Character{
		name:        "Wolverine",
		attackPower: 30,
		defense:     60,
		health:      100,
	})
	fmt.Println(len(testBoss.fighterList.engagedFighters))
	fmt.Println(testBoss.fighterList.engagedFighters[0].name)
	go testBoss.Fight(ch)
	for s := range ch {
		fmt.Println(s)
	}

	testBoss.fighterList.Add(&Character{
		name:        "Captain Marvel",
		attackPower: 25,
		defense:     65,
		health:      110,
	})
	fmt.Println(len(testBoss.fighterList.engagedFighters))

	ch = make(chan string)
	go testBoss.Fight(ch)

	for s := range ch {
		fmt.Println(s)
	}

	for _, v := range majorVillains.list {
		fmt.Printf("Boss Name: %v Capacity: %v\n", v.character.name, v.fighterList.capacity)
	}

	c := make([]chan string, numHeroes)

	fmt.Println("BOSS FIGHT!\n")

	for i, _ := range heroes.list {
		c[i] = make(chan string)
		go EngageWithBoss(i*numBosses/numHeroes, (i+1)*numBosses/numHeroes, heroes.list[i], majorVillains.list, c[i])
	}

	for i := range c {
		for s := range c[i] {
			fmt.Println(s)
		}
	}
}

func main() {
	HeroesVsVillains()
	BossFight()
}

func EngageWithBoss(i, n int, hero *Character, bossList []*BossCharacter, c chan string) {
	defer close(c)

	// Notes / Ideas
	// for each boss
	// engage boss
	// fight until a character is defeated
	// if hero dies return
	// if a boss dies move onto the next boss

	// one go routing should add heroes to boss slices
	// another should iterate over the boss array and fight when the capacity is full
}

func SaveTheWorld(i, n int, hero *Character, villainList []*Character, c chan string) {
	defer close(c)

	for ; i < n; i++ {
		if villainList[i].health > 0 && hero.health > 0 {
			Battle(hero, villainList[i], c)
		}

		if hero.health > 0 && villainList[i].health > 0 {
			Battle(villainList[i], hero, c)
		}
	}
}

func (boss BossCharacter) Fight(c chan string) {
	defer close(c)
	if len(boss.fighterList.engagedFighters) == boss.fighterList.capacity {
		for _, v := range boss.fighterList.engagedFighters {
			Battle(v, boss.character, c)
			Battle(boss.character, v, c)
		}
	} else {
		c <- fmt.Sprint("The fight cannot commence yet.")
	}
}

func (boss BossCharacter) Engage(c *Character) {
	if len(boss.fighterList.engagedFighters) < boss.fighterList.capacity {
		testBoss.fighterList.engagedFighters = append(testBoss.fighterList.engagedFighters, c)
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

type BossCharacter struct {
	character   *Character
	fighterList SafeEngagedList
}

type BossCharacterList struct {
	list []*BossCharacter
}

type CharacterList struct {
	list []*Character
}

var testBoss = BossCharacter{
	&Character{
		name:        "Sabretooth",
		attackPower: 35,
		defense:     55,
		health:      90,
	},
	SafeEngagedList{
		engagedFighters: make([]*Character, 0),
		capacity:        2,
	},
}
var majorVillains = BossCharacterList{
	[]*BossCharacter{
		&BossCharacter{
			&Character{
				name:        "Thanos",
				attackPower: 25,
				defense:     100,
				health:      200,
			},
			SafeEngagedList{
				engagedFighters: make([]*Character, 0),
				capacity:        4,
			},
		},
		&BossCharacter{
			&Character{
				name:        "Ultron",
				attackPower: 15,
				defense:     80,
				health:      150,
			},
			SafeEngagedList{
				engagedFighters: make([]*Character, 0),
				capacity:        4,
			},
		},
	},
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
