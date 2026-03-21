package main
import "fmt"
func main() {
    var idade int
    var nome string
    fmt.Scan(&nome, &idade)

    switch {
    case idade < 12:
        fmt.Println(nome,"eh crianca")
    case idade < 18:
        fmt.Println(nome, "eh jovem")
    case idade < 65:
        fmt.Println(nome, "eh adulto")
    case idade < 100:
        fmt.Println(nome, "eh idoso")
    default:
        fmt.Println(nome, "eh mumia")
    }
}