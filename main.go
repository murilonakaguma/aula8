package main
import (
	"fmt"
)
func main(){	
 votos := make(map[string]int)

 votos["Ana"] += 1
 votos["carlos"]+= 1
 votos["ana"] += 1
 votos["bruno"] += 1
 votos["ana"] += 1
 votos["carlos"] += 1
 for candidato, total := range votos {
	fmt.Printf("%s recebeu %d votos\n", candidato,total)
 }
}

