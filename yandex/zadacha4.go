package main
import (
  "fmt"
)

func main() {
    var n, q int 
    var a, x, z []int
    var b, c, y []int
    fmt.Scan(&n) //1
    var tmp int
    var tmp1 int
    //dohod
    for i := 1; i < n + 1; i++{
        fmt.Scan(&tmp)
        a = append(a, tmp) //2
    }
    //obrazovanie
    for i := 1; i < n + 1; i++{
        fmt.Scan(&tmp1)
        b = append(b, tmp1) //3
    }
    //grj roditelei
    for i := 1; i < n + 1; i++{
        fmt.Scan(&tmp1)
        c = append(c, tmp1) //4
    }
    
    fmt.Scan(&q)
    //dohod
    for i := 1; i < q + 1; i++{
        fmt.Scan(&tmp)
        x = append(x, tmp)
    }
    //obrazovanie
    for i := 1; i < q + 1; i++{
        fmt.Scan(&tmp1)
        y = append(y, tmp1)
    }
    //grj roditelei
    for i := 1; i < q + 1; i++{
        fmt.Scan(&tmp)
        z = append(z, tmp)
    }
    
    var ans []int
    for i := 1; i < q + 1; i++{
        ans = append(ans, 0)
    }
    
    for i := 0; i < q; i++{ //prohod po odnoklassnikam
        for j := 0; j < n ; j++{ //prohod po roditelskomu grajdanstvu
            if (c[j] == 1){
                if(z[i] == j + 1){
                    ans[i] = j + 1
                    break
                }
            }
        }
        if (ans[i] != 0){
            continue
        }
        for j := 0; j < n ; j++{ //prohod po dengam + obr
            if (b[j] == 0){ //ne nujno obrazovanie
                if (x[i] >= a[j]){ //sravnivaem dohod
                    ans[i] = j + 1
                    break
                }
            }
            if (b[j] == 1){ //nujno obrazovanie
                if (y[i] == 1) { //sravnivaem obrazovanie
                    if (x[i] >= a[j]){ //sravnivaem dohod
                        ans[i] = j + 1
                        break
                    }
                }
            }
        }
    }
    for i := 0; i < q; i++ {
        fmt.Printf("%d ", ans[i])
    }
}
