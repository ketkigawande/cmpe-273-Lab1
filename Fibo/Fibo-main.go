package main
import c "ketki/Fibo/funtionality"
import "fmt"

func main(){
var count float64
fmt.Println("Enter n for f(n)")
fmt.Scanln(&count)
if(count<0){
return -1}
answer:= c.Fibo(count)
fmt.Println(answer)
}

