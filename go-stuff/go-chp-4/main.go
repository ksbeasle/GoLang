package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* ----------------------------4.9----------------------------*/

	//Count the frequency of words from a file
	file := "test.txt"
	wf := wordFreq(file)
	fmt.Printf("word\tcount\n")
	for k, v := range wf {
		fmt.Printf("%q\t%d\n", k, v)
	}
	/* ----------------------------4.9----------------------------*/
}

func wordFreq(f string) map[string]int {
	var m = make(map[string]int)
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()

		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	return m
}
