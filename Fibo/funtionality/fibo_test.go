package functionality

import "testing"
func Test_forpositive(t *testing.T) {
var v float64
v = Fibo(10)
if v != 55 {
t.Error("Expected 55, got ", v)
}
}

func Test_forzero(t *testing.T) {
var v float64
v = Fibo(0)
if v != 0 {
t.Error("Expected 0, got ", v)
}
}

func Test_fornegative(t *testing.T) {
var v float64
v = Fibo(-5)
if v>0 {
t.Error("ERROR !!! You entered a negative number ")
}
}






