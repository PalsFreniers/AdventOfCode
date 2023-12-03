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
        getRatio_at(x int, y int) (bool, int)
        extractNumberLeft(x int, y int) ([]Point, string)
        extractNumberRight(x int, y int) ([]Point, string)
}

type Map struct {
        str string
        nbrs []int
        width int
        height int
}

type Point struct {
        x int
        y int
}


func checkPointSlice(sl []Point, pt Point) bool {
        for _, v := range sl {
                if v == pt {
                        return true
                }
        }
        return false
}

func (m Map) extractNumberLeft(x int, y int) ([]Point, string) {
        var checked []Point
        if !isdigit(m.at(x, y)) {
                checked = append(checked, Point{x: x, y: y})
                return checked, ""
        }
        if x > 0 && isdigit(m.at(x - 1, y)) {
                ch, str := m.extractNumberLeft(x - 1, y)
                for _, v := range ch {
                        checked = append(checked, v)
                }
                checked = append(checked, Point{x: x, y: y})
                return checked, str + string([]byte{m.at(x, y)})
        } else {
                checked = append(checked, Point{x: x, y: y})
                return checked, string([]byte{m.at(x, y)})
        }
}

func (m Map) extractNumberRight(x int, y int) ([]Point, string) { 
        var checked []Point
        if !isdigit(m.at(x, y)) {
                checked = append(checked, Point{x: x, y: y})
                return checked, ""
        }
        if x < (m.width - 1) && isdigit(m.at(x + 1, y)) {
                ch, str := m.extractNumberRight(x + 1, y)
                for _, v := range ch {
                        checked = append(checked, v)
                }
                checked = append(checked, Point{x: x, y: y})
                return checked, string([]byte{m.at(x, y)}) + str 
        } else {
                checked = append(checked, Point{x: x, y: y})
                return checked, string([]byte{m.at(x, y)})
        }
}

func (m Map) getRatio_at(x int, y int) (bool, int) {
        if m.at(x, y) != '*' {
                return false, 0
        }
        var numbers []string
        var checked []Point

        if y > 0 {
                yt := -1
                for xt := -1; xt < 2; xt++ {
                        if checkPointSlice(checked, Point{x: x + xt, y: y + yt}) {
                                continue
                        }
                        if isdigit(m.at(x + xt, y + yt)) {
                                ch, fpart := m.extractNumberLeft(x + xt, y + yt)
                                for _, v := range ch {
                                        checked = append(checked, v)
                                }
                                ch, lpart := m.extractNumberRight(x + xt + 1, y + yt)
                                for _, v := range ch {
                                        checked = append(checked, v)
                                }
                                numbers = append(numbers, fpart + lpart)
                        }
                }
        }
        yt := 0
        for xt := -1; xt < 2; xt++ {
                if checkPointSlice(checked, Point{x: x + xt, y: y + yt}) {
                        continue
                }
                if isdigit(m.at(x + xt, y + yt)) {
                        ch, fpart := m.extractNumberLeft(x + xt, y + yt)
                        for _, v := range ch {
                                checked = append(checked, v)
                        }
                        ch, lpart := m.extractNumberRight(x + xt + 1, y + yt)
                        for _, v := range ch {
                                checked = append(checked, v)
                        }
                        numbers = append(numbers, fpart + lpart)
                }
        } 
        if y < (m.height - 1) {
                yt := 1
                for xt := -1; xt < 2; xt++ {
                        if checkPointSlice(checked, Point{x: x + xt, y: y + yt}) {
                                continue
                        }
                        if isdigit(m.at(x + xt, y + yt)) {
                                ch, fpart := m.extractNumberLeft(x + xt, y + yt)
                                for _, v := range ch {
                                        checked = append(checked, v)
                                }
                                ch, lpart := m.extractNumberRight(x + xt + 1, y + yt)
                                for _, v := range ch {
                                        checked = append(checked, v)
                                }
                                numbers = append(numbers, fpart + lpart)
                        }
                }
        }
        
        if len(numbers) != 2 {
                return false, 0
        } else {
                n1, err := strconv.Atoi(numbers[0])
                check(err)
                n2, err := strconv.Atoi(numbers[1])
                check(err)
                return true, (n1 * n2)
        }
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
        if y > 0 {
                if x > 0 && m.at(x - 1, y - 1) != c && !isdigit(m.at(x - 1, y - 1)) {
                        return true;
                }
                if m.at(x, y - 1) != c && !isdigit(m.at(x, y - 1)) { 
                        return true;
                }
                if x < m.width - 1 && m.at(x + 1, y - 1) != c && !isdigit(m.at(x + 1, y - 1)) {
                        return true;
                }
        }
        if x > 0 && m.at(x - 1, y) != c && !isdigit(m.at(x - 1, y)) {
                return true;
        }
        if x < m.width - 1 && m.at(x + 1, y) != c && !isdigit(m.at(x + 1, y)) {
                return true;
        }
        if y < (m.height - 1) {
                if x > 0 && m.at(x - 1, y + 1) != c && !isdigit(m.at(x - 1, y + 1)) {
                        return true;
                }
                if m.at(x, y + 1) != c && !isdigit(m.at(x, y + 1)) { 
                        return true;
                }
                if x < m.width - 1 && m.at(x + 1, y + 1) != c && !isdigit(m.at(x + 1, y + 1)) {
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
        f, err := os.Open("inputs/3.input")
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
                str := strings.TrimSpace(scan.Text())
                m.width = len(str)
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
                        }
                        if isdigit(m.at(x, y)) {
                                ind = true
                                isOk := m.check_around(x, y, '.')
                                if isOk {
                                        comp += m.nbrs[nbIndex]
                                        nbIndex++
                                        ind = false 
                                        for x < m.width && isdigit(m.at(x, y)) {
                                                x++
                                        }
                                }
                        }
                }
        }
        return comp
}

func e3part2() int {
        f, err := os.Open("inputs/3.input")
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
                str := strings.TrimSpace(scan.Text())
                m.width = len(str)
                m.height++
                m.str += str
        }
        
        comp := 0
        for y := 0; y < m.height; y++ {
                for x := 0; x < m.width; x++ {
                        if m.at(x, y) == '*' {
                                isGear, ratio := m.getRatio_at(x, y);
                                if isGear {
                                        comp += ratio
                                }
                        }
                }
        }
        return comp
}

func e3() {
        fmt.Println("Exercice 3 : ")
        fmt.Println("\tPart1 : (", e3part1(), ")")
        fmt.Println("\tPart2 : (", e3part2(), ")")
}
