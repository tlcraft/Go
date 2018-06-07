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
	"sync"
)

var numHeroes = len(heroes.list)
var numVillains = len(villains.list)
var numBosses = len(majorVillains.list)

type SafeIsEngaged struct {
	isEngaged bool
	mux       sync.Mutex
}

func (m *SafeIsEngaged) MarkIsEngaged(e bool) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.isEngaged = e
}

func (m *SafeIsEngaged) IsEngaged() bool {
	m.mux.Lock()
	defer m.mux.Unlock()

	return m.isEngaged
}

type SafeEngagedList struct {
	engagedFighters []*HeroCharacter
	capacity        int
	mux             sync.Mutex
}

func (m *SafeEngagedList) Add(fighter *HeroCharacter) bool {
	m.mux.Lock()
	defer m.mux.Unlock()

	var isAdded bool = false

	if len(m.engagedFighters) < m.capacity {
		m.engagedFighters = append(m.engagedFighters, fighter)
		isAdded = true
	}

	return isAdded
}

func (m *SafeEngagedList) CanFight() bool {
	m.mux.Lock()
	defer m.mux.Unlock()

	var canFight = false
	if m.capacity == len(m.engagedFighters) {
		for _, v := range m.engagedFighters {
			if v.character.health > 0 {
				canFight = true
			} else {
				canFight = false
				break
			}
		}
	}
	return canFight
}

func (b BossCharacterList) Disengage() {
	for _, v := range b.list {
		if v.character.health <= 0 {
			for _, v := range v.fighterList.engagedFighters {
				v.engaged.MarkIsEngaged(false)
			}
		}
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

func Test() {
	fmt.Println("\nTesting out a Boss Character idea")
	ch := make(chan string)

	fmt.Println(testBoss.character.name)
	fmt.Println(len(testBoss.fighterList.engagedFighters))
	testBoss.fighterList.Add(&HeroCharacter{
		&Character{
			name:        "Wolverine",
			attackPower: 30,
			defense:     60,
			health:      100,
		},
		SafeIsEngaged{
			isEngaged: false,
		},
	})

	fmt.Println(len(testBoss.fighterList.engagedFighters))
	fmt.Println(testBoss.fighterList.engagedFighters[0].character.name)
	go testBoss.Fight(ch)
	for s := range ch {
		fmt.Println(s)
	}

	testBoss.fighterList.Add(&HeroCharacter{
		&Character{
			name:        "Captain Marvel",
			attackPower: 25,
			defense:     65,
			health:      110,
		},
		SafeIsEngaged{
			isEngaged: false,
		},
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

	BossesAllDefeatedTest()
	HeroesAllDefeatedTest()
	EndGameBossTest()
	EndGameHeroTest()

	BossesAllDefeatedMultiTest()
	HeroesAllDefeatedMultiTest()
	EndGameHeroMultiTest()
}

func EndGameBossTest() {
	fmt.Println("EndGameBossTest")
	var bossListTest = BossCharacterList{
		[]*BossCharacter{
			&BossCharacter{
				&Character{
					name:        "BossTest",
					attackPower: 25,
					defense:     100,
					health:      200,
				},
				SafeEngagedList{
					engagedFighters: make([]*HeroCharacter, 0),
					capacity:        4,
				},
			},
		},
	}

	var heroListTest = HeroCharacterList{
		[]*HeroCharacter{
			&HeroCharacter{
				&Character{
					name:        "HeroTest",
					attackPower: 20,
					defense:     50,
					health:      70,
				},
				SafeIsEngaged{
					isEngaged: false,
				},
			},
		},
	}

	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
	bossListTest.list[0].character.health = 0
	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
}

func EndGameHeroTest() {
	fmt.Println("EndGameHeroTest")
	var bossListTest = BossCharacterList{
		[]*BossCharacter{
			&BossCharacter{
				&Character{
					name:        "BossTest",
					attackPower: 25,
					defense:     100,
					health:      200,
				},
				SafeEngagedList{
					engagedFighters: make([]*HeroCharacter, 0),
					capacity:        4,
				},
			},
		},
	}

	var heroListTest = HeroCharacterList{
		[]*HeroCharacter{
			&HeroCharacter{
				&Character{
					name:        "HeroTest",
					attackPower: 20,
					defense:     50,
					health:      70,
				},
				SafeIsEngaged{
					isEngaged: false,
				},
			},
		},
	}

	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
	heroListTest.list[0].character.health = 0
	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
}

func EndGameHeroMultiTest() {
	fmt.Println("EndGameHeroMultiTest")
	var bossListTest = BossCharacterList{
		[]*BossCharacter{
			&BossCharacter{
				&Character{
					name:        "BossTest",
					attackPower: 25,
					defense:     100,
					health:      200,
				},
				SafeEngagedList{
					engagedFighters: make([]*HeroCharacter, 0),
					capacity:        4,
				},
			},
			&BossCharacter{
				&Character{
					name:        "BossXTest",
					attackPower: 25,
					defense:     100,
					health:      200,
				},
				SafeEngagedList{
					engagedFighters: make([]*HeroCharacter, 0),
					capacity:        4,
				},
			},
		},
	}

	var heroListTest = HeroCharacterList{
		[]*HeroCharacter{
			&HeroCharacter{
				&Character{
					name:        "HeroTest",
					attackPower: 20,
					defense:     50,
					health:      70,
				},
				SafeIsEngaged{
					isEngaged: false,
				},
			},
		},
	}

	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
	bossListTest.list[1].character.health = 0
	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
	heroListTest.list[0].character.health = 0
	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
	bossListTest.list[0].character.health = 0
	fmt.Println("Is the game over?", EndGame(bossListTest, heroListTest))
}

func BossesAllDefeatedTest() {
	fmt.Println("BossesAllDefeatedTest")
	var bossListTest = BossCharacterList{
		[]*BossCharacter{
			&BossCharacter{
				&Character{
					name:        "BossTest",
					attackPower: 25,
					defense:     100,
					health:      200,
				},
				SafeEngagedList{
					engagedFighters: make([]*HeroCharacter, 0),
					capacity:        4,
				},
			},
		},
	}

	fmt.Println("Are all the bosses defeated?", bossListTest.AllBossesDefeated())
	bossListTest.list[0].character.health = 0
	fmt.Println("Are all the bosses defeated?", bossListTest.AllBossesDefeated())
}

func BossesAllDefeatedMultiTest() {
	fmt.Println("BossesAllDefeatedMultiTest")
	var bossListTest = BossCharacterList{
		[]*BossCharacter{
			&BossCharacter{
				&Character{
					name:        "BossTest",
					attackPower: 25,
					defense:     100,
					health:      200,
				},
				SafeEngagedList{
					engagedFighters: make([]*HeroCharacter, 0),
					capacity:        4,
				},
			},
			&BossCharacter{
				&Character{
					name:        "BossMultiTest",
					attackPower: 25,
					defense:     100,
					health:      200,
				},
				SafeEngagedList{
					engagedFighters: make([]*HeroCharacter, 0),
					capacity:        4,
				},
			},
		},
	}

	fmt.Println("Are all the bosses defeated?", bossListTest.AllBossesDefeated())
	bossListTest.list[0].character.health = 0
	fmt.Println("Are all the bosses defeated?", bossListTest.AllBossesDefeated())
	bossListTest.list[1].character.health = 0
	fmt.Println("Are all the bosses defeated?", bossListTest.AllBossesDefeated())
}

func HeroesAllDefeatedTest() {
	fmt.Println("HeroesAllDefeatedTest")
	var heroListTest = HeroCharacterList{
		[]*HeroCharacter{
			&HeroCharacter{
				&Character{
					name:        "HeroTest",
					attackPower: 20,
					defense:     50,
					health:      70,
				},
				SafeIsEngaged{
					isEngaged: false,
				},
			},
		},
	}

	fmt.Println("Are all the heroes defeated?", heroListTest.AllHeroesDefeated())
	heroListTest.list[0].character.health = 0
	fmt.Println("Are all the heroes defeated?", heroListTest.AllHeroesDefeated())
}

func HeroesAllDefeatedMultiTest() {
	fmt.Println("HeroesAllDefeatedMultiTest")
	var heroListTest = HeroCharacterList{
		[]*HeroCharacter{
			&HeroCharacter{
				&Character{
					name:        "HeroTest",
					attackPower: 20,
					defense:     50,
					health:      70,
				},
				SafeIsEngaged{
					isEngaged: false,
				},
			},
			&HeroCharacter{
				&Character{
					name:        "HeroMultiTest",
					attackPower: 20,
					defense:     50,
					health:      70,
				},
				SafeIsEngaged{
					isEngaged: false,
				},
			},
		},
	}

	fmt.Println("Are all the heroes defeated?", heroListTest.AllHeroesDefeated())
	heroListTest.list[0].character.health = 0
	fmt.Println("Are all the heroes defeated?", heroListTest.AllHeroesDefeated())
	heroListTest.list[1].character.health = 0
	fmt.Println("Are all the heroes defeated?", heroListTest.AllHeroesDefeated())
}

func BossFight() {
	// TODO New Direction
	// Start Go routines to send heroes to fight the bosses
	// Engage heroes with boss characters if the capacity is not full
	// Fight until death and then continue iterating over boss array until one side has won
	// When heroes free up send them on to fight again after they recover some amount of health

	fmt.Println("BOSS FIGHT!\n")
	var count int = 0 // Add a temporary short circuit
	for !EndGame(majorVillains, heroCharacters) && count < 10 {
		var c = make([]chan string, len(majorVillains.list))

		// Assign heroes to villains
		for _, v := range majorVillains.list {
			if v.character.health > 0 {
				for v.fighterList.capacity != len(v.fighterList.engagedFighters) {
					nextHero, err := heroCharacters.NextHero()
					if err == nil {
						v.fighterList.Add(nextHero)
					} else {
						fmt.Println(err)
						break
					}
				}
			}
		}
		count++
		// Fight!
		for i, v := range majorVillains.list {
			c[i] = make(chan string)
			go v.Fight(c[i])
		}

		for i := range c {
			for s := range c[i] {
				fmt.Println(s)
			}
		}

		majorVillains.Disengage()
	}

	fmt.Println("Bosses defeated?", majorVillains.AllBossesDefeated())
	fmt.Println("Heroes defeated?", heroCharacters.AllHeroesDefeated())
}

func main() {
	//HeroesVsVillains()
	//Test()
	BossFight()
}

func EngageWithBoss(i, n int, heroList *HeroCharacterList, bossList []*BossCharacter, c chan string) {
	defer close(c)
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

func PrintStats(c CharacterList, header string) {
	fmt.Println(header)
	for _, v := range c.list {
		fmt.Printf("Name: %v\n\tAttack Power: %v\n\tDefense: %v\n\tHealth: %v\n", v.name, v.attackPower, v.defense, v.health)
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

func (c BossCharacterList) AllBossesDefeated() bool {
	var isDefeated = false
	for _, v := range c.list {
		if v.character.health <= 0 {
			isDefeated = true
		} else {
			isDefeated = false
			break
		}
	}
	return isDefeated
}

func (boss BossCharacter) Fight(c chan string) {
	defer close(c)
	if boss.fighterList.CanFight() {
		for _, v := range boss.fighterList.engagedFighters {
			Battle(v.character, boss.character, c)
			Battle(boss.character, v.character, c)
		}
	} else {
		c <- fmt.Sprint("The fight cannot commence yet.")
	}
}

func (h HeroCharacterList) NextHero() (*HeroCharacter, error) {
	for _, v := range h.list {
		if v.engaged.IsEngaged() == false {
			v.engaged.MarkIsEngaged(true)
			return v, nil
		}
	}

	return &HeroCharacter{}, GameError("No hero is available to fight")
}

func (c HeroCharacterList) AllHeroesDefeated() bool {
	var isDefeated = false
	for _, v := range c.list {
		if v.character.health <= 0 {
			isDefeated = true
		} else {
			isDefeated = false
			break
		}
	}
	return isDefeated
}

func EndGame(b BossCharacterList, h HeroCharacterList) bool {
	if b.AllBossesDefeated() {
		return true
	}

	if h.AllHeroesDefeated() {
		return true
	}

	return false
}

type GameError string

func (e GameError) Error() string {
	return string(e)
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

type HeroCharacter struct {
	character *Character
	engaged   SafeIsEngaged
}

type HeroCharacterList struct {
	list []*HeroCharacter
}

type BossCharacter struct {
	character   *Character
	fighterList SafeEngagedList
}

type BossCharacterList struct {
	list []*BossCharacter
}

var testBoss = BossCharacter{
	&Character{
		name:        "Sabretooth",
		attackPower: 10,
		defense:     55,
		health:      90,
	},
	SafeEngagedList{
		engagedFighters: make([]*HeroCharacter, 0),
		capacity:        1,
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
				engagedFighters: make([]*HeroCharacter, 0),
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
				engagedFighters: make([]*HeroCharacter, 0),
				capacity:        2,
			},
		},
		&BossCharacter{
			&Character{
				name:        "Sabretooth",
				attackPower: 10,
				defense:     55,
				health:      90,
			},
			SafeEngagedList{
				engagedFighters: make([]*HeroCharacter, 0),
				capacity:        1,
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
		&Character{
			name:        "Captain Marvel",
			attackPower: 25,
			defense:     65,
			health:      110,
		},
		&Character{
			name:        "Wolverine",
			attackPower: 30,
			defense:     60,
			health:      100,
		},
	},
}

var heroCharacters = HeroCharacterList{
	[]*HeroCharacter{
		&HeroCharacter{
			&Character{
				name:        "Thor",
				attackPower: 20,
				defense:     50,
				health:      70,
			},
			SafeIsEngaged{
				isEngaged: false,
			},
		},
		&HeroCharacter{
			&Character{
				name:        "Iron Man",
				attackPower: 15,
				defense:     45,
				health:      60,
			},
			SafeIsEngaged{
				isEngaged: false,
			},
		},
		&HeroCharacter{
			&Character{
				name:        "Spider-Man",
				attackPower: 10,
				defense:     30,
				health:      50,
			},
			SafeIsEngaged{
				isEngaged: false,
			},
		},
		&HeroCharacter{
			&Character{
				name:        "Captain Marvel",
				attackPower: 25,
				defense:     65,
				health:      110,
			},
			SafeIsEngaged{
				isEngaged: false,
			},
		},
		&HeroCharacter{
			&Character{
				name:        "Wolverine",
				attackPower: 30,
				defense:     60,
				health:      100,
			},
			SafeIsEngaged{
				isEngaged: false,
			},
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
