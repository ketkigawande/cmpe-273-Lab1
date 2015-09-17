package functionality

func Fibo( n float64 ) (float64){
if(n<0){
return -1}
if(n==1 || n==0){
return n}
y:=Fibo(n-1)+Fibo(n-2)
return y
}
