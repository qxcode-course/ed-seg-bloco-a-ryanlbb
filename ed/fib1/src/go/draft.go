package main
import "fmt"

func coelhosEmMeses(meses, pares int) {
    if meses <= 1 {
        fmt.Println()
        return
    }


    coelhosEmMeses(meses-1, pares)
}

func fib(n int) int{
if n == 0 || n == 1 {
        return 1
    } else { 

        return fib(n-1) + fib(n-2)
    }
    
}

func main() {
    var meses, pares int
    fmt.Scan(&meses, &pares)
    coelhosEmMeses(meses, pares)
}
