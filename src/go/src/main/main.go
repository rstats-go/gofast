package main

/*
  #define USE_RINTERNALS
  #include <R.h>
  #include <Rinternals.h>

*/
import "C"
import "gofast"
import "unsafe"

//export Gofast
func Gofast( x string ) C.SEXP {
  functions := gofast.Gofast(x)
  n := len(functions)

  var out C.SEXP = C.Rf_allocVector( C.STRSXP, C.long(n) )
  C.Rf_protect(out)
  defer C.Rf_unprotect(1)

  for i, s := range functions {
    var cs *C.char = (*C.char)( unsafe.Pointer(& []byte(s) [0]))

    C.SET_STRING_ELT( out, C.R_xlen_t(i), C.Rf_mkChar(cs) )
  }

  return out ;
}

func main() {}
