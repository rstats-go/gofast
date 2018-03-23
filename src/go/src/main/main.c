#include "_cgo_export.h"

SEXP gofast( SEXP x ){
  if( TYPEOF(x) != STRSXP ) error("expecting an string") ;
  SEXP sx = STRING_ELT(x, 0) ;
  GoString gos = { (char*)CHAR(sx), SHORT_VEC_LENGTH(sx) } ;
  GoString res = Gofast(gos) ;

  return ScalarString(mkCharLenCE(res.p, res.n, CE_UTF8)) ;

}
