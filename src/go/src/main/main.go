package main

/*
  #define USE_RINTERNALS
  #include <R.h>
  #include <Rinternals.h>

*/
import "C"
import "gofast"

//export Twice
func Twice( x int32 ) int32 {
  return gofast.Twice(x)
}

func main() {}
