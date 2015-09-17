package main
import "fmt"
import "time"

func main() {
var duration int 
fmt.Println("For how much time(millisecond) do you want make the function sleep???? ")
fmt.Scanln(&duration)
CustomSleep(duration)
}



func CustomSleep(x int){
 
	<- time.After(time.Millisecond * time.Duration(x))
	fmt.Println("TIMED OUT")
	
}


