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

/*
縦 H マス、横 W マスの H × W マスからなる迷路 S があります。
上から i 行目、左から j 列目のマス は S_ij とあらわされ、
S_ij が「#」のとき壁であり、「.」のとき道です。
整数 r、c が与えられるので、S_rc が壁かどうか判定してください。

H W r c
S_1
S_2
...
S_H

1 ≦ H, W ≦ 100
1 ≦ r ≦ H
1 ≦ c ≦ W
S_i は「#」(壁)または「.」(道)からなる W 文字の文字列

S_rc が壁なら「Yes」を、壁ではないなら「No」と出力してください。
*/
