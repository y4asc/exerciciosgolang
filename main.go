package main 

import ( 
    "fmt" 
) 

func main() { 
    var saldo float32
    fmt.Println("Insira o valor inicial:") 
    fmt.Scan(&saldo) 
    escolherOperacao(&saldo) 
} 
func escolherOperacao(saldo *float32) { 
    var opcao int 
    fmt.Println("Você deseja sacar (1), depositar (2) ou encerrar a sessão (3)?") 
    fmt.Scan(&opcao) 
    switch opcao { 
    case 1: 
        sacar(saldo) 
    case 2: 
        depositar(saldo) 
    case 3: 
        encerrar() 
    default: 
        fmt.Println("Opção inválida.") 
    } 
} 
func depositar(saldo *float32) { 
    var valor float32 
    fmt.Println("Qual valor deseja depositar?") 
    fmt.Scan(&valor) 
    *saldo += valor 
    fmt.Println("Seu saldo atual é:", *saldo) 
} 
func sacar(saldo *float32) { 
    var valor float32 
    fmt.Println("Qual valor deseja sacar?") 
    fmt.Scan(&valor) 
    if valor > *saldo { 
        fmt.Println("Saldo insuficiente.") 
    } else { 
        *saldo -= valor 
        fmt.Println("Seu saldo atual é:", *saldo) 
    } 
}  
func encerrar() { 

    fmt.Println("Sessão encerrada.") 
} 