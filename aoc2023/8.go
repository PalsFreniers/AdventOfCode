package main

import (
        "os"
        "fmt"
        "bufio"
        "strings"
)

type path struct {
        left string
        right string
}

func getMovesNumber(base string, moves string, data map[string]path, p2 bool) int {
        str := base
        s := data[str]
        movePos := 0
        comp := 0
        for (str != "ZZZ" && !p2) || (str[2] != 'Z' && p2) {
                if movePos == len(moves) {
                        movePos = 0
                }
                if moves[movePos] == 'L' {
                        str = s.left
                } else if moves[movePos] == 'R' {
                        str = s.right
                }
                movePos++
                comp++
                s = data[str]
        }
        return comp;
}

func e8part1() int {
        f, err := os.Open("inputs/8.input")
        check(err)

        scan := bufio.NewScanner(f)
        scan.Scan()
        moves := scan.Text()
        scan.Scan()

        data := map[string]path{"": {left: "", right: ""}}
        for scan.Scan() {
                str := scan.Text()
                strs := strings.Split(str, " = ")
                strPaths := strings.Split(strings.Trim(strs[1], "()"), ", ")
                data[strs[0]] = path{strPaths[0], strPaths[1]}
        }
        return getMovesNumber("AAA", moves, data, false)
}

func e8part2() int {
        f, err := os.Open("inputs/8.input")
        check(err)

        scan := bufio.NewScanner(f)
        scan.Scan()
        moves := scan.Text()
        scan.Scan()

        data := map[string]path{"": {left: "", right: ""}}
        paths := map[string]int{"": 0}
        for scan.Scan() {
                str := scan.Text()
                strs := strings.Split(str, " = ")
                strPaths := strings.Split(strings.Trim(strs[1], "()"), ", ")
                data[strs[0]] = path{strPaths[0], strPaths[1]}
                if strs[0][2] == 'A' {
                        paths[strs[0]] = 0
                }
        }
        for k := range paths {
                if(k != "") {
                        getMovesNumber(k, moves, data, true)
                }
        }
        return 0
}


func e8() {
        fmt.Println("Exercice 8 : ")
        fmt.Println("\tPart1 : (", e8part1(), ")")
        fmt.Println("\tPart2 : (", e8part2(), ") TODO : LCM")
}
