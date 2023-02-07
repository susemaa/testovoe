package main
import (
  "fmt"
)

func main() {
    var n, x, t, k int
    var a, dif []int
    var aa, min int
    fmt.Scan(&n, &x, &t)
    for i := 0; i < n; i++ {
        fmt.Scan(&aa)
        a = append(a, aa)
        dif = append(dif, Abs(a[i] - x))
        if (dif[i] == 0){
            k++
        }
    }
    for t > 0 {
        min = 1000000000
        for i := 0; i < n; i++ {
            if (dif[i] < min && dif[i] != 0) {
                min = dif[i]
            }
            
        }
        for i := 0; i < n; i++ {
            if (dif[i] == min) {
                dif[i] -= min
                t -= min
                k++
                if (t < 0){
                    dif[i] += min
                    k--
                }
            }
        }
    }
    fmt.Println(k)
    for i := 0; i < n; i++ {
        if(dif[i] == 0){
            fmt.Printf("%d ",i+1)
        }
    }
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
