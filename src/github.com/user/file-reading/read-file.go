package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
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
     if regex == REGEX_3 {
        fmt.Println("range_m")
        fmt.Println(info[1])
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
        fmt.Println(line)
        info, regex := parseLine(line)
        fillContext(info, regex, context)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
