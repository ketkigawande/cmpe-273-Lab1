package main

import "testing"
import "time"
func Test_sleep(t *testing.T) {
v:=time.Now()
CustomSleep(800)
w:=time.Now()
if v==w {
t.Error("Expected different times, got same time after calling sleep function ")
}
}



