package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

var REGEX_1 = "There is a ([a-zA-Z ]+) at position ([0-9]+)"

var REGEX_2 = "([a-zA-Z ]+) attack is ([0-9]+)"

var REGEX_3 = "Resources are ([0-9]+) meters away"

var REGEX_4 = "([a-zA-Z ]+) has ([0-9]+) hp"

var REGEX_5 = "([a-zA-Z ]+) is Enemy"

func logErr(err error) {
     if err != nil {
         log.Fatal(err)
     }
}

type Character struct {
     hp int
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
     range_m int
     enemy_map map[Enemy]int
}

type Context struct {
     hero Hero
     field Field
     enemy_map map[string]Enemy
}

func isStringMatches(line,regex string) bool {
    r,err := regexp.Compile(regex)
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

func parseLine(line string) ([]string,string) {
    regex := whichRegexIsAppropiate(line)
    r, err := regexp.Compile(regex)
    logErr(err)
    result := r.FindStringSubmatch(line)
    fmt.Println(result)
    return result, regex
}

func fillContext(info []string, regex string, context *Context){
     fmt.Println()
     fmt.Println(regex)
     if regex == REGEX_1 {
        enemy := info[1]
        position := info[2]
        fmt.Printf("%q %q\n",enemy, position)
     } else if regex == REGEX_2 {
        character := info[1]
        attackPoint := info[2]
        fmt.Printf("%q %q\n", character, attackPoint)
     } else if regex == REGEX_3 {
        fmt.Printf("%q\n",info[1])
     } else if regex == REGEX_4 {
        character := info[1]
        hp := info[2]
        hpInt, err := strconv.Atoi(hp)
        if err == nil {
          hpInt = 0
        }
        fmt.Printf("%q %q\n", character,hp)
        if character == "Hero" {
           herotemp := context.hero
           if (Hero{}) == herotemp {
              herotemp = Hero{Character{hp:hpInt}}
           } else {
              herotemp.hp = hpInt
           }
           context.hero = herotemp
        } else {
           enemytemp ,err := context.enemy_map[character]
           if !err {
              enemytemp = Enemy{}
           }
           enemytemp.species = character
           enemytemp.hp = hpInt
        }
     } else if regex == REGEX_5 {
        species := info[1]
        fmt.Printf("%q %q\n",species)
        enemytemp,err := context.enemy_map[species]
        if !err {
           enemytemp = Enemy{}
        }
        enemytemp.species = species
        context.enemy_map[species] = enemytemp
     }
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

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
