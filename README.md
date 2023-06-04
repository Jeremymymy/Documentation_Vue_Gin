# Documentation

### 開發階段 <br>
```    git clone https://github.com/Jeremymymy/Documentation_Vue_Gin.git  ``` <br> <br>

安裝 Node.js 18.16.0 以上的版本 "https://nodejs.org/en" <br>
安裝 Vue-Cli 2.9.6 以上的版本: <br>
```    npm install -g vue-cli  ``` <br> <br>
安裝 Quasar-Cli 2.0.1 以上的版本: <br>
```    npm install -g quasar-cli@2.0.1  ``` <br> <br>

安裝 Golang 1.20 版，可於網頁安裝 "https://go.dev/dl/" <br>
安裝 Gin 框架: <br>
```    go get github.com/gin-gonic/gin  ``` <br> <br>

安裝 MySQL，建立一個名為 tsmcdocs 的 database <br>
若 database 的帳號非 root 或密碼非 root，需至 dbconnect/DBconnect.go 更改 dsn <br>
    dsn := "你的MySQL帳號:你的MySQL密碼@tcp(127.0.0.1:3306)/tsmcdocs?charset=utf8mb4&parseTime=True&loc=Local" <br> <br>

啟動後端程式: <br>
``` cd Documentation_Vue_Gin/ ``` <br>
``` go run main.go ``` <br> <br>

啟動前端程式: <br>
``` cd Documentation_Vue_Gin/quasar-project/ ``` <br>
``` npm install ``` <br>
``` quasar dev ``` <br> <br>

### 佈署階段 <br>
啟用本機 Hyper-V <br>
安裝 WSL2 ，可參考 "https://learn.microsoft.com/zh-tw/windows/wsl/install" <br>
安裝 Docker Desktop ，可參考 "https://dockerdocs.cn/docker-for-windows/install/" <br>
至 dbconnect/DBconnect.go 更改dsn ，將原先連線至本機 MySQL 的 dsn 註解，把下方 dsn 取消註解<br>
    dsn := "root:root@tcp(db)/tsmcdocs?charset=utf8mb4&parseTime=True&loc=Local" <br>
``` cd Documentation_Vue_Gin/ ``` <br>
``` mkdir mysql_data ``` <br> 
``` docker-compose up ``` <br>