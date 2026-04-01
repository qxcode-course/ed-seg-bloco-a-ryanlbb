package main
import "fmt"
func main() {
    var n, rotacoes int
    fmt.Scan(&n, &rotacoes)
    vet := make([]int, n)
    rotacionado := make([]int, n)
    for i := range n {
        vet[i] = i + 1
        if rotacoes > 0 {
            rotacionado[(i + rotacoes) % n] = i + 1 
            continue
        }
        rotacionado[i] = i + 1
    }

    fmt.Print("[ ")
    for _, v := range rotacionado {
        fmt.Print(v, " ")
    }
    fmt.Println("]")
}
