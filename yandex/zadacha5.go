package main //yandex1
import (
	"fmt"
)

func main() {
	var n, k int
	var s, alph string
	var p, d, m []int
	alph = "abcdefghijklmnopqrstuvwxyz"
	fmt.Scan(&n, &k)
	fmt.Scan(&s)
	var t int
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		p = append(p, t)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		d = append(d, t)
	}
	for i := 0; i < n; i++ {
		m = append(m, 0)
	}
	var value, zxc int
	var temp byte
	var resstr string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[j] = 0
		}
		resstr = ""
		temp = s[i]
		for j := 0; j < k; j++ {
			num := gnfl(temp, s)
			if m[num] == 0 { //perviiy primer m[a] == m[0], m[b] == m[1], m[z] == m[2]
				m[num]++
			} else {
				for ij := 0; ij < 26; ij++ {
					if temp == alph[ij] {
						zxc = ij
						break
					}
				}
				m[num]++
				temp = alph[(zxc+(m[num]-1)*d[num])%26]
			}
			temp1 := string(temp)
			resstr += temp1
			value += countvalue(resstr)
			temp = s[p[num]-1]
		}
	}
	fmt.Println(value)
}

func countvalue(s string) int {
	var al string
	for i := 0; i < len(s); i++ {
		flag := 0
		for j := 0; j < len(al); j++ {
			if s[i] == al[j] {
				flag = 1
			}
		}
		if flag == 0 {
			s1 := string(s[i])
			al += s1
		}
	}
	return len(al)
}

func gnfl(temp byte, s string) int { //get number from letter
	for i := 0; i < len(s); i++ {
		if temp == s[i] {
			return i
		}
	}
	return -1
}
