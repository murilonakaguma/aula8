package main
import (
	"fmt"
)
func main(){	
capitais := map[string]string{
	"SP" : "São paulo",
	"RJ" : "Rio de Janeiro",
	"ES" : "Espirito Santo",
}
   capitais["BH"] = "Belo Horizonte"
   for k, v := range capitais {
	fmt.Println("sigla,",k," nome",  v)
   }
   delete(capitais, "ES")
   for k,v:= range capitais {
	fmt.Println("sigla",k ,"nome", v)
   }
}

