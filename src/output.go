package src

import (
	"fmt"
	"strings"
)

func ReadOut(s []string, our_string string, termlen int, pos string) {
	if len(our_string) > 0 {
		e := map[rune][]string{}
		var b string
		var q rune = 32
		count := 0
		for i := 1; i < len(s); i += 9 {
			e[q] = s[i : i+8]
			q++
		}
		for _, v := range our_string {
			if string(v) == "\\" {
				count++
			}
		}
		if count*2 == len(our_string) {
			for i := 0; i < count; i++ {
				fmt.Println()
			}
			return
		}
		g := 167
		k := strings.ReplaceAll(our_string, "\\n", "\n")
		l := strings.Split(k, "\n")
		strg := []string{}
		for _, w := range l {
			for i := 0; i < 8; i++ {
				for t := 0; t < len(w); t++ {
					if w[t] >= 32 && w[t] <= 126 {
						b += e[rune(w[t])][i]
					}
				}
				strg = append(strg, b)
				b = ""
			}
		}
		kil := ""
		for i := 0; i < 8; i++ {
			for j := 0; j < g; j++ {
				if j == g/2-len(strg[i])/2 {
					for re := 0; re < len(strg[i]); re++ {
						kil += string(strg[i][re])
					}
				}
				if j == g/2-len(strg[i])/2 {
					kil += "\n"
					j = g
				} else {
					kil += " "
				}
			}
		}
		fmt.Print(kil)
	}
}
