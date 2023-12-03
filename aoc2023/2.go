package main

import (
        "fmt"
        "os"
        "bufio"
        "strings"
        "strconv"
)

func e2part1() int {
        f, err := os.Open("2.input")
        check(err)
        defer f.Close()
        
        scan := bufio.NewScanner(f)
        scan.Split(bufio.ScanLines)

        gameNb := 1
        comp := 0
        for scan.Scan() {
                str := scan.Text()
                i := 0
                for str[i] != ':' {
                        i++
                }
                i++
                str = str[i:len(str)]
                game := strings.Split(str, ";")
                isValid := true
                for _, v := range game {
                        r := 0
                        g := 0
                        b := 0
                        batch := strings.Split(v, " ")
                        var nb int
                        for _, ba := range batch {
                                if len(ba) == 0 {
                                        
                                } else if isdigit(ba[0]) {
                                        nb, err = strconv.Atoi(ba)
                                        check(err)
                                } else {
                                        if len(ba) >= 3 && ba[0:3] == "red" {
                                                r += nb
                                        } else if len(ba) >= 5 && ba[0:5] == "green" {
                                                g += nb
                                        } else if len(ba) >= 4 && ba[0:4] == "blue" {
                                                b += nb
                                        }
                                }
                        }
                        if !(r <= 12 && g <= 13 && b <= 14) {
                                isValid = false;
                                break
                        }
                }
                if isValid {
                        comp += gameNb
                }
                gameNb++
        }
        return comp
}

func e2part2() int {
        f, err := os.Open("2.input")
        check(err)
        defer f.Close()
        
        scan := bufio.NewScanner(f)
        scan.Split(bufio.ScanLines)

        gameNb := 1
        comp := 0
        for scan.Scan() {
                str := scan.Text()
                i := 0
                for str[i] != ':' {
                        i++
                }
                i++
                str = str[i:len(str)]
                game := strings.Split(str, ";")
                r := 0
                g := 0
                b := 0
                for _, v := range game {
                        batch := strings.Split(v, " ")
                        var nb int
                        for _, ba := range batch {
                                if len(ba) == 0 {
                                        
                                } else if isdigit(ba[0]) {
                                        nb, err = strconv.Atoi(ba)
                                        check(err)
                                } else {
                                        if len(ba) >= 3 && ba[0:3] == "red" {
                                                if nb > r {
                                                        r = nb
                                                }
                                        } else if len(ba) >= 5 && ba[0:5] == "green" {
                                                if nb > g {
                                                        g = nb
                                                }
                                        } else if len(ba) >= 4 && ba[0:4] == "blue" {
                                                if nb > b {
                                                        b = nb
                                                }
                                        }
                                }
                        }
                }
                comp += r * g * b
                gameNb++
        }
        return comp
}

func e2() {
        fmt.Println("Exercise 2 : ")
        fmt.Println("\tPart1 : (", e2part1(), ")")
        fmt.Println("\tPart2 : (", e2part2(), ")")
}

