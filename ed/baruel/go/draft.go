package main
import "fmt"
func main() {
    var totalFig, nBaruelFig int 
    fmt.Scan(&totalFig)
    figs := make([]int, totalFig)
    figsBaruel := make(map[int]int)
    repetidas := make([]int, 0)

    // todas as fig
    for i := range totalFig {
        figs[i] = i + 1
    }

    fmt.Scan(&nBaruelFig)

    // fig que o caba tem
    for range nBaruelFig {
        var a int
        fmt.Scan(&a)
        figsBaruel[a]++

        if figsBaruel[a] > 1 {
            repetidas = append(repetidas, a)
        }
    }

    for i, v := range figsBaruel {
        if v == figsBaruel[i] {
            v = 0
        } 
    }

    if len(repetidas) == 0 {
        fmt.Println("N")
    } else {
        for _, v := range repetidas {
            fmt.Print(v, " ")
        }
    }
    fmt.Println()

    fmt.Println(figs)
}