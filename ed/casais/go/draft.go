package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    animais :=  make(map[int]int)
    casados := 0

    for range n {
        var a int
        fmt.Scan(&a)
        animais[a]++

        if animais[-a] > 0 {
            casados++
            animais[-a]--
            animais[a]--
        } 
    }

    fmt.Println(casados)
}