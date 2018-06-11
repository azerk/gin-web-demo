package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"log"
	"io"
	"fmt"
	"gin_web/utils"
	"path/filepath"
)

func main() {
	isSuccess:=true
	dirs := []string{"temp"} // 设置需要释放的目录

	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := asset.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
	}
	gin.SetMode(gin.ReleaseMode)
	router:=gin.Default()
	router.GET("/user", func(context *gin.Context) {
		name:=context.Query("name")
		age:=context.DefaultQuery("age","0")
		context.String(http.StatusOK,"%s, %s",name,age)
	})
	router.LoadHTMLGlob("temp/*")
	router.GET("/load", func(context *gin.Context) {
		context.HTML(http.StatusOK,"upload.html",gin.H{})
	})
	router.POST("/uploadPic", func(context *gin.Context) {
		fmt.Println("开始上传图片······")
		file,header,err:=context.Request.FormFile("file")
		if err!=nil {
			fmt.Println("bad request")
			context.String(http.StatusBadRequest,"bad request")
			return
		}
		out,err_file:=os.Create(header.Filename)
		if err_file !=nil{
			fmt.Println("can't create")
			log.Fatal(err)
		}
		defer out.Close()
		_,err_writer:=io.Copy(out,file)
		if err_writer !=nil{
			fmt.Println("can't writer")
			context.String(http.StatusBadRequest,"fail")
			log.Fatal(err)
		}
		context.String(http.StatusOK,"success file")

	})
	router.Run(":8000")
}
