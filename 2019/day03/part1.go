package main

import (
	"fmt"
	"strconv"
)

type WirePath struct {
	Direction byte
	Distance  int
}

type Point struct {
	X int
	Y int
}

func part1(in <-chan string) {
	c := make(chan *WirePath, 10)
	go func() {
		for s := range in {
			var wp WirePath
			wp.Direction = s[0]
			if n, err := strconv.Atoi(s[1:]); err == nil {
				wp.Distance = n
				c <- &wp
			}
		}

		close(c)
	}()

	for wp := range c {
		fmt.Printf("WirePath: <%q,%v>\n", wp.Direction, wp.Distance)
	}

	//wires := make(map[Point]int)
}
