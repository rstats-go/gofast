#' @useDynLib gofast
#' @export
twice <- function(x) {
  .Call("twice", x, PACKAGE = "gofast")
}
