package main

import (
	"fmt"
)

func main() {
	var n, x int
	var p []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		p = append(p, x)
	}
	var pos [6]int
	var amount1, amount2, money1, money2, moneytmp float32
	for i := 0; i < n-1; i++ {
		moneytmp = 1
		amount1 = buy(moneytmp, p[i])
		for j := i + 1; j < n; j++ {
			moneytmp = sell(amount1, p[j])
			if moneytmp > money1 {
				money1 = moneytmp
				pos[0] = i
				pos[1] = j
			}
		}
	}
	for i := 0; i < n-3; i++ {
		moneytmp = 1
		amount2 = buy(moneytmp, p[i])
		for j := i + 1; j < n-2; j++ {
			moneytmp1 := sell(amount2, p[j])
			for ii := j + 1; ii < n-1; ii++ {
				amount21 := buy(moneytmp1, p[ii])
				for jj := ii + 1; jj < n; jj++ {
					moneytmp2 := sell(amount21, p[jj])
					if moneytmp2 > money2 {
						money2 = moneytmp2
						pos[2] = i
						pos[3] = j
						pos[4] = ii
						pos[5] = jj
					}
				}
			}
		}
	}
	if money1 > money2 {
		fmt.Printf("1\n%d %d", pos[0]+1, pos[1]+1)
	} else {
		fmt.Printf("2\n%d %d\n%d %d", pos[2]+1, pos[3]+1, pos[4]+1, pos[5]+1)
	}
}

func buy(money float32, price int) float32 {
	amount := money / float32(price)
	return amount
}

func sell(amount float32, price int) float32 {
	money := amount * float32(price)
	return money
}
