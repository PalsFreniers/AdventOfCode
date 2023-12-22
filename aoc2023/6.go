package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allAtoi(nums []string) []int {
        var ret []int
        for _, v := range nums {
                if v == "" {
                        continue
                }
                n, err := strconv.Atoi(v);
                check(err)
                ret = append(ret, n)
        }
        return ret
}

func e6part1() int {
        f, err := os.Open("inputs/6.input")
        check(err)

        scan := bufio.NewScanner(f)
        scan.Scan()
        time := allAtoi(strings.Split(scan.Text(), " ")[1:])
        scan.Scan()
        distance := allAtoi(strings.Split(scan.Text(), " ")[1:])
        comp := 1
        for i := 0; i < len(time); i++ {
                n := 0
                for speed := 0; speed < time[i]; speed++ {
                        dist := speed * (time[i] - speed)
                        if dist > distance[i] {
                                n++
                        }
                } 
                comp *= n
        }
        return comp
}

func delSpace(num string) int {
        a := ""
        for _, v := range num {
                if v != ' ' {
                        a += string(v)
                }
        }
        dat, err := strconv.Atoi(a)
        check(err)
        return dat
}

func e6part2() int {
        f, err := os.Open("inputs/6.input")
        check(err)

        scan := bufio.NewScanner(f)
        scan.Scan()
        time := delSpace(strings.Split(scan.Text(), ":")[1])
        scan.Scan()
        distance := delSpace(strings.Split(scan.Text(), ":")[1])
        n := 0
        for speed := 0; speed < time; speed++ {
                dist := speed * (time - speed)
                if dist > distance {
                        n++
                }
        } 
        return n
}

func e6() {
        fmt.Println("Exercise 6 : ")
        fmt.Println("\tPart1 : (", e6part1(), ")")
        fmt.Println("\tPart2 : (", e6part2(), ")")
}
