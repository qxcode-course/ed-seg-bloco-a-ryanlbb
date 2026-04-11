package main
import "fmt"

func coelhosEmMeses(meses, pares, mes int) {
    coelhos := 1 
    if meses <= 2 {
        fmt.Println(1)
        return
    }
    if mes == meses {
        fmt.Println(coelhos)
        return
    }

    coelhosEmMeses(meses, pares + coelhos, mes+1)
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
    // coelhosEmMeses(meses, pares, 1)
    fmt.Println(fib(4))
}
