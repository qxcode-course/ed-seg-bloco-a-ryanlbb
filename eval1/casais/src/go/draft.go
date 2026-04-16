package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    animais := make([]int, n)
    animaisMap := make(map[int]int)
    for i := range animais {
        fmt.Scan(&animais[i])
        animaisMap[animais[i]]++
    }

    casais := 0
    for _, v := range animais {
        if animaisMap[v] >= 1 && animaisMap[-v] >= 1 {
            casais++
            animaisMap[v]--
            animaisMap[-v]--
            } 
        }

    fmt.Println(casais)
}

