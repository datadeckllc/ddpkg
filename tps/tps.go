package tps

import (
  "fmt"
  "net/http"
  "log"
  "golang.org/x/net/html"
)

const TPS_URL := "https://truepeoplesearch.com/"

func tps_search(fname, lname, mname string){
  var tps_str string = html.EscapeString(fname + " " +  mname + " " +lname)
  res, e := http.Get(TPS_URL + "results?=" + tps_str)
  e != nil {
    log.Fatal(e)
  }

  defer tps.Body.Close()
  doc, e := html.Parse(res)
  fmt.Printf("%T\n", )
}
