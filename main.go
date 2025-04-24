package main

import "fmt"

func main() {

    estoque := map[string]int{
        "Coxinha":      10,
        "Pão de Queijo": 15,
        "Refrigerante": 20,
    }

    fmt.Println("Estoque inicial:")
    for produto, quantidade := range estoque {
        fmt.Printf("%s: %d\n", produto, quantidade)
    }

    estoque["Coxinha"] -= 2
    estoque["Pão de Queijo"] -= 1

    fmt.Println("\nEstoque atualizado:")
    for produto, quantidade := range estoque {
        fmt.Printf("%s: %d\n", produto, quantidade)
    }
  }