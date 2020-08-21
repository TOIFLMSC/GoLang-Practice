package main

import (
	"fmt"
	"sort"
	"strings"
)

const targetstring = `The Nvidia GeForce RTX 3090 is now confirmed as the next halo graphics card from Team Green, 
thanks to Micron's inadvertent posting of memory details (the PDF is now removed). With that piece of knowledge, 
we've dissected the rest of what we expect to find in the RTX 3090. Nvidia has a countdown to the 21st anniversary 
of its first GPU, the GeForce 256, slated for September 1. The battle for the best 
graphics cards and top of the GPU hierarchy is about to get heated.
We've talked about the Nvidia Ampere and RTX 30-series as a whole elsewhere, so this discussion 
is focused purely on the GeForce RTX 3090. Let's dig into the details of what we know about the 
GeForce RTX 3090, including the expected GPU and memory specifications, release date, price, features, and more.
First, the GeForce RTX 3090 branding is the first 90-series suffix we've seen since the GTX 690 
back in 2012. That was a dual-GPU variant of the GTX 680, but based on the Micron documentation, 
RTX 3090 will still be a single GPU. Spoiler: multi-GPU support in games is practically dead, at 
least on life support. Why bring back the 90 branding? Simple: It opens the door for a new tier 
of performance and pricing. That's not good news for our wallets.
We discussed the Micron inadvertent posting of details and more in a recent Tom's Hardware show, 
which you can view below. Let's dig into the details.`

func main() {
	SplitterFunc := func(r rune) bool { //Splitter function is a function that uses comparator, to compare every rune in target string to set of escape-runes
		return strings.ContainsRune(".,?!:; \n", r)
	}

	SplittedText := strings.FieldsFunc(targetstring, SplitterFunc) //Now, using FieldsFunc function, that separates string to array of strings, that in that case uses our function SplitterFunc to compare runes

	for idx, word := range SplittedText {
		SplittedText[idx] = strings.ToLower(word)
	}

	dir := map[string]int{}
	for _, word := range SplittedText { // Now, in cycle we are comparing all strings from our last stage to our dictionary, if some line is not in the dictionary we add it into it, if it is present, we increase our counter
		if _, ok := dir[word]; !ok {
			dir[word] = 1
		} else if _, ok := dir[word]; ok {
			dir[word] = dir[word] + 1
		}
	}

	var val []int
	for _, valx := range dir {
		val = append(val, valx)
	}

	sort.Ints(val)

	for i := 1; i <= 10; i++ {
		for key, valx := range dir {
			if valx == val[len(val)-i] {
				fmt.Printf("Word: '%s' has %d occurrences.\n", key, valx)
				delete(dir, key)
				break
			}
		}
	}

}
