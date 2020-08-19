package part1

import (
	"log"
	"os"
	"strconv"
)

func unpackEx(s1 string) (s2 string) {
	s2 = ""
	if s1 == "" {
		os.Exit(1)
	}
	var last byte
	var escapecheck bool
	for i := 0; i < len(s1); i++ {
		if byte('a') <= s1[i] && s1[i] <= byte('z') {
			//fmt.Printf("%c is letter\n", s1[i])
			last = s1[i]
			s2 = s2 + string(last)
		} else if byte('0') <= s1[i] && s1[i] <= byte('9') && last == byte('\\') && escapecheck == true {
			//fmt.Printf("%c is number after double escape\n", s1[i])
			a, err := strconv.Atoi(string(s1[i]))
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < a-1; i++ {
				s2 = s2 + string(last)
			}
			last = s1[i]
		} else if byte('0') <= s1[i] && s1[i] <= byte('9') && last == byte('\\') {
			//fmt.Printf("%c is number after escape\n", s1[i])
			last = s1[i]
			s2 = s2 + string(last)
		} else if (byte('0') <= s1[i] && s1[i] <= byte('9')) && (byte('0') <= last && last <= byte('9')) {
			//fmt.Printf("%c is number that must be written %c\n", last, s1[i])
			a, err := strconv.Atoi(string(s1[i]))
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < a-1; i++ {
				s2 = s2 + string(last)
			}
			last = s1[i]
		} else if byte('\\') == s1[i] && last == byte('\\') {
			//fmt.Printf("%c is escape doubled symbol\n", s1[i])
			escapecheck = true
			last = s1[i]
			s2 = s2 + string(last)
		} else if byte('0') <= s1[i] && s1[i] <= byte('9') {
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
		} else if byte('\\') == s1[i] {
			//fmt.Printf("%c is escape symbol\n", s1[i])
			last = s1[i]
			escapecheck = false
		}
	}
	return s2
}
