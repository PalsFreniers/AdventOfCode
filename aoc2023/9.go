package main

import (
        "os"
        "fmt"
        "bufio"
        "strings"
        "strconv"
)

func atoiTab(tab []string) []int {
        var ret []int
        for _, v := range tab {
                i, err := strconv.Atoi(v)
                check(err)
                ret = append(ret, i)
        }
        return ret
}

type Layer []int

func (l Layer) generateNextLayer() Layer {
        var ret Layer
        for i := 0; i < len(l) - 1; i++ {
                ret = append(ret, l[i + 1] - l[i])
        }
        return ret
}

func (l Layer) checkLastLayer() bool {
        for _, v := range l {
                if v != 0 {
                        return false
                }
        }
        return true
}

func (l Layer) Print() {
        fmt.Printf("{ ")
        for _, v := range l {
                fmt.Printf("%d, ", v)
        }
        fmt.Printf("}\n")
}

func e9part1() int {
        f, err := os.Open("inputs/9.input")
        check(err)

        scan := bufio.NewScanner(f)
        comp := 0
        for scan.Scan() {
                str := scan.Text()
                var tree []Layer
                numbers := Layer(atoiTab(strings.Split(str, " ")))
                for !(numbers.checkLastLayer()) {
                        tree = append(tree, numbers)
                        numbers = numbers.generateNextLayer()
                }
                tree = append(tree, numbers)
                tree[len(tree) - 1] = append(tree[len(tree) - 1], 0)
                for i := len(tree) - 2; i >= 0; i-- {
                        ldata := tree[i][len(tree[i]) - 1]
                        abs := tree[i + 1][len(tree[i + 1]) - 1]
                        newHistory := ldata + abs
                        tree[i] = append(tree[i], newHistory)
                }
                comp += tree[0][len(tree[0]) - 1]
        }
        return comp
}

func e9part2() int {
        f, err := os.Open("inputs/9.input")
        check(err)

        scan := bufio.NewScanner(f)
        comp := 0
        for scan.Scan() {
                str := scan.Text()
                var tree []Layer
                numbers := Layer(atoiTab(strings.Split(str, " ")))
                for !(numbers.checkLastLayer()) {
                        tree = append(tree, numbers)
                        numbers = numbers.generateNextLayer()
                }
                tree = append(tree, numbers)
                tree[len(tree) - 1] = append(Layer([]int{0}), tree[len(tree) - 1]...)
                for i := len(tree) - 2; i >= 0; i-- {
                        ldata := tree[i][0]
                        abs := tree[i + 1][0]
                        newHistory := ldata - abs
                        tree[i] = append(Layer([]int{newHistory}), tree[i]...)
                }
                comp += tree[0][0]
        }
        return comp
}

func e9() {
        fmt.Println("Exercice 9 : ")
        fmt.Println("\tPart1 : (", e9part1(), ")")
        fmt.Println("\tPart2 : (", e9part2(), ")")
}
