package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"output/src"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 4 {
		fmt.Println()
		str := os.Args[1]
		banner := os.Args[2]
		flag := os.Args[3]

		if src.Checkflag(flag) == nil {
			q := strings.Split(flag, "=")

			if src.Validstr(str, banner) == nil {
				if src.Banner(banner) == nil {
					file, err := os.Open(banner + ".txt")
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()
					scanner := bufio.NewScanner(file)
					ascii := []string{}
					for scanner.Scan() {
						s := strings.ReplaceAll(scanner.Text(), "\n", "")
						ascii = append(ascii, s)
					}
					if err := scanner.Err(); err != nil {
						log.Fatal(err)
					}
					termlen := consoleSize()
					src.ReadOut(ascii, str, termlen, q[1])

				} else {
					fmt.Println(src.Banner(banner))
				}
			} else {
				fmt.Println(src.Validstr(str, banner))
			}
		} else {
			fmt.Println(src.Checkflag(flag))
		}
	}
}

func consoleSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}
