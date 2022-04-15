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

	var A []int
	var a int
	var str string
	sc.Scan()
	s = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		a, _ = strconv.Atoi(s[i])
		A = append(A, a)
	}

	for i := 0; i < q; i++ {
		sc.Scan()
		switch sc.Text() {
		case "1":
			A = A[:len(A)-1]

		case "2":
			str = fmt.Sprintf("%v", A)
			str = strings.Trim(str, "[]")
			fmt.Println(str)

		default:
			s = strings.Split(sc.Text(), " ")
			a, _ = strconv.Atoi(s[1])
			A = append(A, a)

		}
	}
}

/*
入力
N Q
A_1 A_2 A_3 ... A_N
query_1
query_2
...
query_Q

1 ≦ N ≦ 100
1 ≦ Q ≦ 100
1 ≦ A_i ≦ 100 (1 ≦ i ≦ N)
query_i は 0 x(末尾x追加) または 1(pop_back:末尾要素削除) または 2(全要素出力)
1 ≦ x ≦ 100
pop_back 操作が指定される際、Aの要素数は 1 以上
*/
