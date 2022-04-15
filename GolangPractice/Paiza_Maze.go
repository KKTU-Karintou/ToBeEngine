package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	var maze [][]string
	var s = strings.Split(sc.Text(), " ")
	var h, _ = strconv.Atoi(s[0])
	var r, _ = strconv.Atoi(s[2])
	var c, _ = strconv.Atoi(s[3])

	for i := 0; i < h; i++ {
		sc.Scan()
		maze = append(maze, strings.Split(sc.Text(), ""))
	}

	if maze[r-1][c-1] == "." {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
