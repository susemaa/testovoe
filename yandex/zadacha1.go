package main
import (
  "fmt"
)

func main() {
    var n, k, x int
    var symbs, rows, essay []int
    //s4itaem kol-vo bukv
    fmt.Scan(&n)
    //s4itaem vse bukvi
    for i := 0; i < n; i++{
        fmt.Scan(&x)
        symbs = append(symbs, x)
    }
    //s4itaem ryadi bukv
    for i := 0; i < n; i++{
        fmt.Scan(&x)
        rows = append(rows, x)
    }
    //s4itaem kol-vo bukv v essay
    fmt.Scan(&k)
    //s4titaem essay
    for i := 0; i < k; i++{
        fmt.Scan(&x)
        essay = append(essay, x)
    }
    
    var ans, num1, num2 int
    for i := 0; i < k - 1; i++ {
        for j := 0; j < n; j++ {
            if (essay[i] == symbs[j]) {
                num1 = j
                break
            }
        }
        for j := 0; j < n; j++ {
            if (essay[i + 1] == symbs[j]) {
                num2 = j
                break
            }
        }
        if (rows[num1] != rows[num2]) {
            ans++
        }
    }
    
    fmt.Println(ans);
}
