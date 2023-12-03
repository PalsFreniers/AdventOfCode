package main

import (
        "fmt"
        "os"
        "bufio"
        "strings"
        "strconv"
)

type Map_func interface {
        at(x int, y int) char
        check_at(x int, y int, c char) bool
}

type Map struct {
        str string
        nbrs []int
        width int
        height int
}

func (m Map) at(x int, y int) char {
        if(x < 0 || x >= m.width) {
                return 0;
        }
        if(y < 0 || y >= m.height) {
                return 0;
        }
        return m.str[x + y * m.width];
}

func (m Map) check_around(x int, y int, c char) bool {
        fmt.Printf("Checked at [%3d; %3d] : { ", x, y)
        defer fmt.Println("}")
        if y > 0 {
                if x > 0 {
                        fmt.Printf("%c, ", m.at(x - 1, y - 1))
                }
                if x > 0 && m.at(x - 1, y - 1) != c && !isdigit(m.at(x - 1, y - 1)) {
                        return true;
                }
                fmt.Printf("%c, ", m.at(x, y - 1))
                if m.at(x, y - 1) != c && !isdigit(m.at(x, y - 1)) { 
                        return true;
                }
                if x < m.width {
                        fmt.Printf("%c, ", m.at(x + 1, y - 1))
                }
                if x < m.width && m.at(x + 1, y - 1) != c && !isdigit(m.at(x + 1, y - 1)) {
                        return true;
                }
        }
        if x > 0 {
                fmt.Printf("%c, ", m.at(x - 1, y))
        }
        if x > 0 && m.at(x - 1, y) != c && !isdigit(m.at(x - 1, y)) {
                return true;
        }
        fmt.Printf("%c, ", m.at(x, y))
        if x < m.width {
                fmt.Printf("%c, ", m.at(x + 1, y))
        }
        if x < m.width && m.at(x + 1, y) != c && !isdigit(m.at(x + 1, y)) {
                return true;
        }
        if y != (m.height - 1) {
                if x > 0 {
                        fmt.Printf("%c, ", m.at(x - 1, y + 1))
                }
                if x > 0 && m.at(x - 1, y + 1) != c && !isdigit(m.at(x - 1, y + 1)) {
                        return true;
                }
                fmt.Printf("%c, ", m.at(x, y + 1))
                if m.at(x, y + 1) != c && !isdigit(m.at(x, y + 1)) { 
                        return true;
                }
                if x < m.width {
                        fmt.Printf("%c, ", m.at(x + 1, y + 1))
                }
                if x < m.width && m.at(x + 1, y + 1) != c && !isdigit(m.at(x + 1, y + 1)) {
                        return true;
                }
        }
        return false
}

func mapAtoi(str string) int {
        ret, err := strconv.Atoi(str)
        check(err)
        return ret
}

func ft_strmap(strs []string) []int {
        ret := []int{}
        for _, v := range strs {
                ret = append(ret, mapAtoi(v))
        }
        return ret
}

func e3part1() int {
        f, err := os.Open("3.input")
        check(err)
        defer f.Close()
        
        scan := bufio.NewScanner(f)
        scan.Split(bufio.ScanLines)
        
        m := Map {
                str: "",
                nbrs: nil,
                width: 0,
                height: 0,
        }

        for scan.Scan() {
                str := strings.TrimRight(scan.Text(), "\r\n")
                if m.width == 0 {
                        m.width = len(str)
                }
                m.height++
                m.str += str
        }

        m.nbrs = ft_strmap(strings.FieldsFunc(m.str, func(r rune) bool { return !isdigitr(r) }))

        nbIndex := 0
        comp := 0
        ind := false
        for y := 0; y < m.height; y++ {
                for x := 0; x < m.width; x++ {
                        if !isdigit(m.at(x, y)) && ind {
                                nbIndex++
                                ind = false 
                                fmt.Println("comp is still :", comp, "\n")
                        }
                        if isdigit(m.at(x, y)) {
                                ind = true
                                isOk := m.check_around(x, y, '.')
                                fmt.Printf("[%t] => %d\n", isOk, m.nbrs[nbIndex])
                                if isOk {
                                        comp += m.nbrs[nbIndex]
                                        nbIndex++
                                        ind = false 
                                        fmt.Println("comp is now :", comp, "\n")
                                        for x < m.width && isdigit(m.at(x, y)) {
                                                x++
                                        }
                                }
                        }
                }
        }
        fmt.Println("{", nbIndex, "; ", len(m.nbrs), "}")
        return comp
}

func e3() {
        fmt.Println("Exercice 3 : ")
        fmt.Println("\tPart1 : (", e3part1(), ")")
//         fmt.Println("\tPart2 : (", e3part2(), ")")
}
