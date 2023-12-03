package main

import (
        "fmt"
        "os"
        "bufio"
)

func e1part1() int {
        f, err := os.Open("1.input")
        check(err)
        defer f.Close()

        scan := bufio.NewScanner(f)
        scan.Split(bufio.ScanLines)

        hasFirst := false
        comp := 0
        for scan.Scan() {
                str := scan.Text()
                first  := 10
                last := 10
                for i := 0; i < len(str); i++ {
                        if isdigit(str[i]) {
                                if hasFirst == false {
                                        first = int(str[i] - '0')
                                        hasFirst = true
                                } else {
                                        last = int(str[i] - '0')
                                }
                        }
                }
                hasFirst = false
                if last == 10 {
                        last = first
                }
                tmp := (first * 10) + last
                comp += tmp
        }
        return comp;
}

func e1part2() int {
        f, err := os.Open("1.input")
        check(err)
        defer f.Close()

        scan := bufio.NewScanner(f)
        scan.Split(bufio.ScanLines)

        hasFirst := false
        comp := 0
        t := 0
        for scan.Scan() {
                str := scan.Text()
                first := 0
                last := 0
                lstr := len(str)
                for i := 0; i < lstr; i++ {
                        if isdigit(str[i]) {
                                if hasFirst == false {
                                        first = int(str[i] - '0')
                                        hasFirst = true
                                } else {
                                        last = int(str[i] - '0')
                                }
                        } else if str[i] == 'o' && i + 3 <= lstr && str[i:i+3] == "one" {
                                if hasFirst == false {
                                        first = 1
                                        hasFirst = true
                                } else {
                                        last = 1
                                }
                                i++
                        } else if str[i] == 't' && i + 3 <= lstr && str[i:i+3] == "two" {
                                if hasFirst == false {
                                        first = 2
                                        hasFirst = true
                                } else {
                                        last = 2
                                }
                                i++
                        } else if str[i] == 't' && i + 5 <= lstr && str[i:i+5] == "three" {
                                if hasFirst == false {
                                        first = 3
                                        hasFirst = true
                                } else {
                                        last = 3
                                }
                                i++
                        } else if str[i] == 'f' && i + 4 <= lstr && str[i:i+4] == "four" {
                                if hasFirst == false {
                                        first = 4
                                        hasFirst = true
                                } else {
                                        last = 4
                                }
                                i++
                        } else if str[i] == 'f' && i + 4 <= lstr && str[i:i+4] == "five" {
                                if hasFirst == false {
                                        first = 5
                                        hasFirst = true
                                } else {
                                        last = 5
                                }
                                i++
                        } else if str[i] == 's' && i + 3 <= lstr && str[i:i+3] == "six" {
                                if hasFirst == false {
                                        first = 6
                                        hasFirst = true
                                } else {
                                        last = 6
                                }
                                i++
                        } else if str[i] == 's' && i + 5 <= lstr && str[i:i+5] == "seven" {
                                if hasFirst == false {
                                        first = 7
                                        hasFirst = true
                                } else {
                                        last = 7
                                }
                                i++
                        } else if str[i] == 'e' && i + 5 <= lstr && str[i:i+5] == "eight" {
                                if hasFirst == false {
                                        first = 8
                                        hasFirst = true
                                } else {
                                        last = 8
                                }
                                i++
                        } else if str[i] == 'n' && i + 4 <= lstr && str[i:i+4] == "nine" {
                                if hasFirst == false {
                                        first = 9
                                        hasFirst = true
                                } else {
                                        last = 9
                                }
                                i++
                        }
                }
                hasFirst = false
                if last == 0 {
                        last = first
                }
                t++
                tmp := (first * 10) + last
                comp += tmp
        }
        return comp;
}

func e1() {
        fmt.Println("Exercise 1 :")
        fmt.Println("\tPart 1 : (", e1part1(), ")")
        fmt.Println("\tPart 2 : (", e1part2(), ")")
}
