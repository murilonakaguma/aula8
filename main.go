package main

import (
    "fmt"
)

func main() {
        estoque := map[string]int{
            "COXINHA":       10,
            "PÃO DE QUEIJO": 15,
            "REFRIGERANTE":  20,
        }
    
    
        estoque["COXINHA"] -= 2
        estoque["PÃO DE QUEIJO"] -= 1
    
    
        fmt.Println("Estoque atual:")
        for produto, quantidade := range estoque {
            fmt.Printf("%s: %d unidades\n", produto, quantidade)
        }
    }

