package main
import "fmt"
func main() {
    var n_consulta, n_buscas int
    fmt.Scan(&n_consulta)
    consulta := make(map[string]int)
    for range n_consulta {
        var coisa string  
        fmt.Scan(&coisa)
        consulta[coisa]++
    }

    fmt.Scan(&n_buscas)
    buscas := make([]string, n_buscas)
    matchs := make([]int, n_buscas) 
    for i := range n_buscas {
        fmt.Scan(&buscas[i])
        matchs[i] = 0
    }

    for i := range buscas {
        if consulta[buscas[i]] > 0 {
            matchs[i] = consulta[buscas[i]]
        }
    }

    for i, v := range matchs {
        fmt.Print(v)
        if i < len(matchs) - 1 {
            fmt.Printf(" ")
        }
    }
    fmt.Println()
}
