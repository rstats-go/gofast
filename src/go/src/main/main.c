#include "_cgo_export.h"

SEXP twice( SEXP x ){
  if( TYPEOF(x) != INTSXP ) error("expecting an integer vector") ;
  return Twice( INTEGER(x)[0] ) ;
}
