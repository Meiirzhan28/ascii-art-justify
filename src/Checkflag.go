package src

import (
	"errors"
	"regexp"
	"strings"
)

func Checkflag(s string) error {
	check := regexp.MustCompile(`(--align=)\d*`)
	if check.MatchString(s) {
		a := strings.Split(s, "=")
		check2 := regexp.MustCompile(`center|right|left|justify`)
		if check2.MatchString(a[1]) {
			return nil
		}
	}
	return errors.New("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --align=right")
}
