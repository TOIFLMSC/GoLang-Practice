package part1

import (
	"log"
	"os"
	"strconv"
)

func unpack(s1 string) (s2 string) {
	s2 = ""
	if s1 == "" {
		os.Exit(1)
	}
	var last byte
	for i := 0; i < len(s1); i++ {
		if byte('a') <= s1[i] && s1[i] <= byte('z') {
			//fmt.Printf("%c is letter\n", s1[i])
			last = s1[i]
			s2 = s2 + string(last)
		}
		if byte('0') <= s1[i] && s1[i] <= byte('9') {
			//fmt.Printf("%c is number\n", s1[i])
			if byte('a') <= last && last <= byte('z') {
				a, err := strconv.Atoi(string(s1[i]))
				if err != nil {
					log.Fatal(err)
				}
				for i := 0; i < a-1; i++ {
					s2 = s2 + string(last)
				}
			}
			last = s1[i]
		}
	}
	return s2
}
