package main

import "fmt"

func main() {
  
  estoque := map[string]int{
        "COXINHA":       10,
        "PÃO DE QUEIJO": 15,
        "REFRIGERANTE":  20,
    }


    estoque["COXINHA"] -= 2    
    estoque["PÃO DE QUEIJO"]-- 

 
    fmt.Println("ESTOQUE ATUALIZADO:")
    for produto, quantidade := range estoque {
        fmt.Printf("%s: %d\n", produto, quantidade)
    }
}