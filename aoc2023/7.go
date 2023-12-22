package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type handType int

const (
        HIGHT_CARDS handType = iota
        ONE_PAIR handType = iota
        TWO_PAIR handType = iota
        THREE_OF_KIND handType = iota
        FULL_HOUSE handType = iota
        FOUR_OF_KIND handType = iota
        FIVE_OF_KIND handType = iota
)

type hand struct {
        cards string
        htype handType
        bid int
}

func initCards() map[char]int {
        cards := map[char]int {
                'A': 0,
                'K': 0,
                'Q': 0,
                'J': 0,
                'T': 0,
                '9': 0,
                '8': 0,
                '7': 0,
                '6': 0,
                '5': 0,
                '4': 0,
                '3': 0,
                '2': 0,
        }
        return cards
}

func initMedian() map[handType]int {
        median := map[handType]int{
                HIGHT_CARDS: 0,
                ONE_PAIR: 0,
                TWO_PAIR: 0,
                THREE_OF_KIND: 0,
                FULL_HOUSE: 0,
                FOUR_OF_KIND: 0,
                FIVE_OF_KIND: 0,
        }
        return median
}

func checkMedian(median map[handType]int) handType {
        if median[FIVE_OF_KIND] != 0 {
                return FIVE_OF_KIND
        }
        if median[FOUR_OF_KIND] != 0 {
                return FOUR_OF_KIND
        }
        if median[THREE_OF_KIND] != 0 {
                if median[ONE_PAIR] != 0 {
                        return FULL_HOUSE
                }
                return THREE_OF_KIND
        }
        if median[ONE_PAIR] == 2 {
                return TWO_PAIR
        }
        if median[ONE_PAIR] == 1 {
                return ONE_PAIR
        }
        return HIGHT_CARDS
}

func getHandType(hnd []char) handType {
        cards := initCards()
        for _, v := range hnd {
                cards[v]++
        }
        median := initMedian()
        for _, v := range cards {
                if v == 2 {
                        median[ONE_PAIR]++
                }
                if v == 3 {
                        median[THREE_OF_KIND]++
                }
                if v == 4 {
                        median[FOUR_OF_KIND]++
                }
                if v == 5 {
                        median[FIVE_OF_KIND]++
                }
        }
        return checkMedian(median)
}

func addAll(hand map[handType]int, cards map[char]int, index char) map[handType]int {
        if hand[FIVE_OF_KIND] != 0 {
                return hand
        }
        if hand[FOUR_OF_KIND] != 0 {
                if cards[index] != 0 {
                        hand[FOUR_OF_KIND] = 0
                        hand[FIVE_OF_KIND] = 1
                }
                return hand
        }
        if hand[THREE_OF_KIND] != 0 {
                if cards[index] == 1 {
                        hand[THREE_OF_KIND] = 0
                        hand[FOUR_OF_KIND] = 1
                }
                if cards[index] == 2 {
                        hand[THREE_OF_KIND] = 0
                        hand[FIVE_OF_KIND] = 1
                }
                if cards[index] == 3 {
                        if hand[ONE_PAIR] != 0 {
                                hand[THREE_OF_KIND] = 0
                                hand[FIVE_OF_KIND] = 1
                                return hand
                        }
                        hand[THREE_OF_KIND] = 0
                        hand[FOUR_OF_KIND] = 1
                }
                return hand
        }
        if hand[ONE_PAIR] == 2 {
                if cards[index] == 1 {
                        hand[ONE_PAIR] = 1
                        hand[THREE_OF_KIND] = 1
                }
                if cards[index] == 2 {
                        hand[ONE_PAIR] = 0
                        hand[FOUR_OF_KIND] = 1
                }
                return hand
        }
        if hand[ONE_PAIR] == 1 {
                if cards[index] != 0 {
                        hand[ONE_PAIR] = 0
                        hand[THREE_OF_KIND] = 1
                }
                return hand
        }
        if cards[index] != 0 {
                hand[ONE_PAIR] = 1
        }
        return hand
}

func getHandTypeJokers(hnd []char) handType {
        cards := initCards()
        for _, v := range hnd {
                cards[v]++
        }
        median := initMedian()
        for _, v := range cards {
                if v == 2 {
                        median[ONE_PAIR]++
                }
                if v == 3 {
                        median[THREE_OF_KIND]++
                }
                if v == 4 {
                        median[FOUR_OF_KIND]++
                }
                if v == 5 {
                        median[FIVE_OF_KIND]++
                }
        }
        median = addAll(median, cards, 'J')
        return checkMedian(median)
}

func needSort(a char, b char) bool {
        c := map[char]int {
                'A': 0,
                'K': 1,
                'Q': 2,
                'J': 3,
                'T': 4,
                '9': 5,
                '8': 6,
                '7': 7,
                '6': 8,
                '5': 9,
                '4': 10,
                '3': 11,
                '2': 12,
        }
        return c[a] > c[b]
}

func needSortJoker(a char, b char) bool {
        c := map[char]int {
                'A': 0,
                'K': 1,
                'Q': 2,
                'T': 3,
                '9': 4,
                '8': 5,
                '7': 6,
                '6': 7,
                '5': 8,
                '4': 9,
                '3': 10,
                '2': 11,
                'J': 12,
        }
        return c[a] > c[b]
}

func e7part1() int {
        f, err := os.Open("inputs/7.input")
        check(err)

        scan := bufio.NewScanner(f)
        var data []hand
        for scan.Scan() {
                str := scan.Text()
                player := strings.Split(str, " ")
                playerHand := player[0]
                playerBid, err := strconv.Atoi(player[1])
                check(err)
                data = append(data, hand{cards: playerHand, htype: getHandType([]char(playerHand)), bid: playerBid})
        }
        sort.Slice(data, func(i int, j int) bool {
                if data[i].htype == data[j].htype {
                        for k, v := range data[i].cards {
                                if byte(v) == data[j].cards[k] {
                                        continue
                                }
                                return needSort(byte(v), data[j].cards[k])
                        }
                        return false
                } else {
                        return data[i].htype < data[j].htype
                }
        })
        comp := 0
        for i := 0; i < len(data); i++ {
                comp += (i + 1) * data[i].bid
        }
        return comp
}

func e7part2() int {
        f, err := os.Open("inputs/7.input")
        check(err)

        scan := bufio.NewScanner(f)
        var data []hand
        for scan.Scan() {
                str := scan.Text()
                player := strings.Split(str, " ")
                playerHand := player[0]
                playerBid, err := strconv.Atoi(player[1])
                check(err)
                data = append(data, hand{cards: playerHand, htype: getHandTypeJokers([]char(playerHand)), bid: playerBid})
        }
        sort.Slice(data, func(i int, j int) bool {
                if data[i].htype == data[j].htype {
                        for k, v := range data[i].cards {
                                if byte(v) == data[j].cards[k] {
                                        continue
                                }
                                return needSortJoker(byte(v), data[j].cards[k])
                        }
                        return false
                } else {
                        return data[i].htype < data[j].htype
                }
        })
        comp := 0
        for i := 0; i < len(data); i++ {
                comp += (i + 1) * data[i].bid
        }
        return comp
}

func e7() {
        fmt.Println("Exercice 7 : ")
        fmt.Println("\tPart1 : (", e7part1(), ")")
        fmt.Println("\tPart2 : (", e7part2(), ")")
}
