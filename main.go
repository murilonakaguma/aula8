package main
import (
	"fmt"
)
func analisarNota(nota1, nota2 float64)(float64, string){ 
	media := (nota1 + nota2) / 2
	var resultado string
	if (media <= 6){
		resultado = "reprovado"
	}else {
		resultado = "aprovado"
	}
	return media, resultado
}
func main(){
	
media, resultado := analisarNota(7.5, 8.5)
fmt.Println("media:", media, "resultado é:", resultado)
}

