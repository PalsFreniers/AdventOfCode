package main

import (
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
        source int
        destination int
        _range int
}

type ruleMap struct {
        name string
        rules []rule
}

func genRule(line string) rule {
        nums := remap(strings.Split(line, " "), func(_ int, val string) int {
                ret, err := strconv.Atoi(val)
                check(err)
                return ret
        })
        ret := rule{
                source: nums[1],
                destination: nums[0],
                _range: nums[2],
        }
        return ret
}

func e5part1() int {
        return 0
}

func e5part2() int {
        return 0
}

func e5() {
        fmt.Println("Exercise 6 : ")
        fmt.Println("\tPart1 : (", e5part1(), ")")
        fmt.Println("\tPart2 : (", e5part2(), ")")
}
