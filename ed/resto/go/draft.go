package main
import "fmt"

func empilhamento(x int) {
    if x <= 0 {
        return
    }

    resto := x % 2
    x = x / 2

    empilhamento(x)
    fmt.Println(x, resto)
}

func main() {
    var n int
    fmt.Scan(&n)
    empilhamento(n)
}
