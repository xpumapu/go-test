package main

import "net/http"
import "io"


func hello(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
    res,
    `<DOCTYPE html>
<html>
  <head>
      <title>heeeeeehhehe World</title>
  </head>
  <body>
      Hello World!
  </body>
</html>`,
  )
}




func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":3000", nil)
    
}