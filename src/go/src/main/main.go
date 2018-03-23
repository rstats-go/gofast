package main

/*
  #define USE_RINTERNALS
  #include <R.h>
  #include <Rinternals.h>

*/
import "C"
import "gofast"

//export Gofast
func Gofast( x string ) string {
  return gofast.Gofast(x)
}

func main() {}
