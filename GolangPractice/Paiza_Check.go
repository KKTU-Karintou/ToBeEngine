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
	var s = strings.Split(sc.Text(), " ")
	var n, _ = strconv.Atoi(s[0])
	var q, _ = strconv.Atoi(s[1])
	var T, Q []string

	for i := 0; i < n; i++ {
		sc.Scan()
		T = append(T, sc.Text())
	}
	for i := 0; i < q; i++ {
		sc.Scan()
		Q = append(Q, sc.Text())
	}

	for i := 0; i < q; i++ {
		for j := 0; j < n; j++ {
			if strings.Compare(T[j], Q[i]) == 0 {
				fmt.Println(j + 1)
				break
			} else if j == n-1 {
				fmt.Println("-1")
			}
		}
	}
}

/*
N 個の文字列 S_1, ... , S_N と、Q 個の文字列 T_1, ... , T_Q が与えられます。
各 T_i について、以下の処理を行ってください。

・ S_j == T_i を満たす最小の j を出力する。
・ ただし、そのような j が存在しない場合は -1 を出力する。

入力
N Q
S_1
S_2
...
S_N
T_1
T_2
...
T_Q

1 ≦ N ≦ 100
1 ≦ Q ≦ 100
S_i, T_j は英小文字からなる1文字以上3文字以下の文字列
				(1 ≦ i ≦ N,1 ≦ j ≦ Q)

Q 行出力してください。
i 行目には、S_j == T_i を満たす最小の j を出力してください。
ただし、そのような j が存在しない場合は -1 を出力してください。
j_1
...
j_Q
*/
