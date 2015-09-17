package main

import "testing"
func Test_interface(t *testing.T) {
var v float64
r:=rectangle {3,4}
s:=shape(r)
v = s.perimeter()
if v != 14 {
t.Error("Expected 14, got ", v)
}
c:=circle{5}
s=shape(c)
v=s.perimeter()
if v!=31.419999999999998 {
t.Error("Expected 31.42,got",v)
}
}


