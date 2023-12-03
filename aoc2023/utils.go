package main

import "unicode"

type char = byte

func check(e error) {
        if e != nil {
                panic(e);
        }
}

func isdigit(c char) bool {
        return (c >= '0' && c <= '9');
}

func isdigitr(c rune) bool {
        return (unicode.IsDigit(c));
}
