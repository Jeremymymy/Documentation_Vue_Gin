# Documentation

安裝 Golang 1.20 版以上，可於網頁安裝 <br>
安裝 Gin 框架: <br>
```    go get github.com/gin-gonic/gin  ``` <br> <br>
安裝 MySQL，建立一個叫做 tsmcdocs 的database <br>
若database的帳號非root 或密碼非root，需至dbconnect/DBconnect.go更改dsn <br>
    dsn := "你的MySQL帳號:你的MySQL密碼@tcp(127.0.0.1:3306)/tsmcdocs?charset=utf8mb4&parseTime=True&loc=Local" <br> <br>
啟動後端程式碼: <br>
``` cd Cloud_DevOps/ ``` <br>
``` go run main.go ``` <br>