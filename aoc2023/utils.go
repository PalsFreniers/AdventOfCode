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

func remap[K any, V any](slc []K, f func(int, K) V) []V {
        var ret []V
        for k, v := range slc {
                ret = append(ret, f(k, v))
        }
        return ret
}
