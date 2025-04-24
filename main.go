package main
import (
	"fmt"
)
func pegarNome()(string, string){ 
  return "barry", "Allen"
}
func main(){
nome, sobrenome := pegarNome()
fmt.Println(nome, sobrenome)
}

