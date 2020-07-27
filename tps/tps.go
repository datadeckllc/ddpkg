package tps

import (
  "fmt"
  "net/http"
  "log"
  "golang.org/x/net/html"
  "io"
)
const TPS_URL = "https://truepeoplesearch.com/"

func Search(name, city, state string) *html.Node {
  var tps_str string = html.EscapeString(name + "&citystatezip=" + city + state)

  res, e := http.Get(TPS_URL + "results?=" + tps_str)
  if e != nil {
    log.Fatalln(e)
  }
  defer res.Body.Close()


  doc, e := html.Parse(io.Reader(res.Body))
  if e != nil {
    log.Fatalln(e)
  }

  return doc
}
