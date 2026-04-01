package main
import "fmt"

func empilhamento(x int) (string) {
    if x != 0 {
        
    }
    div := x / 2
    resto := x  % 2
    fmt.Println(resto)
    return string(div)+ " " +string(resto)
}

func main() {
    var n int
    fmt.Scan(&n)
    empilhamento(n)
}
