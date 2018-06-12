# gin-web-demo
upload pic

使用gin做的上传文件的小demo

包括三个接口： 

1.访问web页面的接口：http://59.110.214.75/load 

2.restful post方式的上传文件的接口：http://59.110.214.75/upload 

3.常用的get方式的访问接口：http://59.110.214.75/user?name=kzz&age=18 

go项目中静态资源无法直接打包，此处使用的是go-bindata,使用方法:

1.下载安装go-bindata:go get -u github.com/jteeuwen/go-bindata/...

2.将静态资源打包（demo中主要为temp文件下）：

  go-bindata -o=gin_web/utils/asset.go -pkg=asset temp/...  
  
  此处会生成asset.go  
  
3.项目启动时释放静态资源到指定目录：  

  此项目是在main.go直接释放：  
  
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
  
