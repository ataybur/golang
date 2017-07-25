package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	REGEX_1  = "There is a ([a-zA-Z ]+) at position ([0-9]+)"
	REGEX_2  = "([a-zA-Z ]+) attack is ([0-9]+)"
	REGEX_3  = "Resources are ([0-9]+) meters away"
	REGEX_4  = "([a-zA-Z ]+) has ([0-9]+) hp"
	REGEX_5  = "([a-zA-Z ]+) is Enemy"
	END_LINE = "\n"
	CONST_1  = "Hero started journey with %d HP!" + END_LINE
	CONST_2  = "Hero defeated %s with %d HP remaining" + END_LINE
	CONST_3  = "Survived" + END_LINE
	CONST_4  = "%s defeated Hero with %d HP remaining" + END_LINE
	CONST_5  = "Hero is Dead!! Last seen at position %d!!" + END_LINE
)

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Character struct {
	hp          int
	attackPoint int
}

type Hero struct {
	Character
}

type Enemy struct {
	Character
	species string
}

type Field struct {
	range_m   int
	enemy_map map[int]Enemy
}

type Context struct {
	hero      Hero
	field     Field
	enemy_map map[string]Enemy
}

func isStringMatches(line, regex string) bool {
	r, err := regexp.Compile(regex)
	logErr(err)
	result := r.FindStringSubmatch(line)
	return len(result) != 0
}

func whichRegexIsAppropiate(line string) string {
	result := ""
	if isStringMatches(line, REGEX_1) {
		result = REGEX_1
	} else if isStringMatches(line, REGEX_2) {
		result = REGEX_2
	} else if isStringMatches(line, REGEX_3) {
		result = REGEX_3
	} else if isStringMatches(line, REGEX_4) {
		result = REGEX_4
	} else if isStringMatches(line, REGEX_5) {
		result = REGEX_5
	}
	return result
}

func parseLine(line string) ([]string, string) {
	regex := whichRegexIsAppropiate(line)
	r, err := regexp.Compile(regex)
	logErr(err)
	result := r.FindStringSubmatch(line)
	fmt.Println(result)
	return result, regex
}

func fillContext(info []string, regex string, context *Context) {
	fmt.Println()
	fmt.Println(regex)
	if regex == REGEX_1 {
		enemy := info[1]
		position := info[2]
		positionInt, err := strconv.Atoi(position)
		if err != nil {
			positionInt = 0
		}
		if len(context.enemy_map) == 0 {
			context.enemy_map = make(map[string]Enemy)
		}
		enemytemp, err2 := context.enemy_map[enemy]
		if err2 {
			enemytemp = Enemy{}
		}
		enemytemp.species = enemy
		if len(context.field.enemy_map) == 0 {
			context.field.enemy_map = make(map[int]Enemy)
		}
		context.field.enemy_map[positionInt] = enemytemp
		fmt.Printf("%q %q\n", enemy, position)
	} else if regex == REGEX_2 {
		character := info[1]
		attackPoint := info[2]
		fmt.Printf("%q %q\n", character, attackPoint)
		attackPointInt, err := strconv.Atoi(info[2])
		if err != nil {
			attackPointInt = 0
		}
		if character == "Hero" {
			context.hero.attackPoint = attackPointInt
		} else {
			enemyTemp := context.enemy_map[character]
			enemyTemp.attackPoint = attackPointInt
			context.enemy_map[character] = enemyTemp
		}
	} else if regex == REGEX_3 {
		fmt.Printf("%q\n", info[1])
		rangeInt, err := strconv.Atoi(info[1])
		if err != nil {
			rangeInt = 0
		}
		context.field.range_m = rangeInt
	} else if regex == REGEX_4 {
		character := info[1]
		hp := info[2]
		hpInt, err := strconv.Atoi(hp)
		if err != nil {
			hpInt = 0
		}
		fmt.Printf("%q %q\n", character, hp)
		if character == "Hero" {
			herotemp := context.hero
			if (Hero{}) == herotemp {
				herotemp = Hero{Character{hp: hpInt}}
			} else {
				herotemp.hp = hpInt
			}
			context.hero = herotemp
		} else {
			enemytemp, ok := context.enemy_map[character]
			if !ok {
				enemytemp = Enemy{}
			}
			enemytemp.species = character
			enemytemp.hp = hpInt
			context.enemy_map[character] = enemytemp
		}
	} else if regex == REGEX_5 {
		species := info[1]
		fmt.Printf("%q\n", species)
		if len(context.enemy_map) == 0 {
			context.enemy_map = make(map[string]Enemy)
		}
		enemytemp, ok := context.enemy_map[species]
		if !ok {
			enemytemp = Enemy{}
		}
		enemytemp.species = species
		context.enemy_map[species] = enemytemp
	}
}

func fight(hero *Hero, enemy Enemy) bool {
	result := false
	heroAttackP := hero.attackPoint
	enemyAttackP := enemy.attackPoint
	enemyHP := enemy.hp
	heroHP := hero.hp
	remains := enemyHP % heroAttackP
	if remains != 0 {
		remains -= heroAttackP
	}
	newEnemyHP := enemyHP + remains
	multiplier := newEnemyHP / heroAttackP
	multipliedEnemyAP := multiplier * enemyAttackP
	enemyName := enemy.species
	if heroHP > multipliedEnemyAP {
		heroHP -= multipliedEnemyAP
		hero.hp = heroHP
		fmt.Printf(CONST_2, enemyName, heroHP)
		result = true
	} else {
		remains := heroHP % enemyAttackP
		if remains != 0 {
			remains -= enemyAttackP
		}
		newHeroHP := heroHP + remains
		multiplier := newHeroHP / enemyAttackP
		multipliedHeroAP := multiplier * heroAttackP
		fmt.Printf(CONST_4, enemyName, enemyHP-multipliedHeroAP)
	}
	return result
}

func main() {
	file, err := os.Open("/home/ataybur/go-workspace/src/github.com/user/file-reading/lines")
	logErr(err)
	defer file.Close()
	var context = new(Context)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		info, regex := parseLine(line)
		fillContext(info, regex, context)
	}
	field := context.field
	isHeroAlive := true
	fmt.Printf(CONST_1, context.hero.hp)
	fmt.Printf("Range is %d"+END_LINE, field.range_m)
	var lastIndex int
	for i := 1; i <= field.range_m; i++ {
		enemy, ok := field.enemy_map[i]
		if ok {
			enemy2, ok2 := context.enemy_map[enemy.species]
			if ok2 {
				isHeroAlive = fight(&context.hero, enemy2)
				if !isHeroAlive {
					lastIndex = i
					break
				}
			}
		}
	}
	if isHeroAlive {
		fmt.Println(CONST_3)
	} else {
		fmt.Printf(CONST_5, lastIndex)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
