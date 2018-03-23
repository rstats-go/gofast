#' @useDynLib gofast
#' @export
gofast <- function(x) {
  .Call("gofast", x, PACKAGE = "gofast")
}
