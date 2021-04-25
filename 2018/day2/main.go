package main

import (
	"bufio"
	"os"
)

func main() {
	in := make(chan string)
	go readLines("./input", in)
	for s := range in {

}

func readLines(path string, out chan<- string) error {
	defer close(out)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		out <- s.Text()
	}
	return nil
}
