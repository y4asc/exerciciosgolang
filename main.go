
package main

import "fmt"

func main() {
    var numeros [5]int

    fmt.Println("Digite 5 números inteiros (pressione Enter após cada um):")

    for i := 0; i < 5; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&numeros[i])
    }

    soma := 0
    for _, num := range numeros {
        soma += num
    }

    fmt.Printf("\nA soma dos números é: %d\n", soma)
    
    fmt.Println("Pressione Enter para sair...")
    fmt.Scanln()
}