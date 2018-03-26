package main

import "github.com/gin-gonic/gin"
import "github.com/buger/jsonparser"
import "net/http"
import "io/ioutil"

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "hello",
    })
  })

  r.GET("/btcprice", func(c *gin.Context) {
    resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
    if err != nil {
      c.JSON(500, gin.H{
        "message": err,
      })  
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    jsonparser.Get(body, "time")

    c.JSON(200, string(body))
  })

  r.Run() // listen and serve on 0.0.0.0:8080
}
