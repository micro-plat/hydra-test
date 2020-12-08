package main

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var app = hydra.NewApp(
	hydra.WithDebug(),
	hydra.WithServerTypes(http.Web),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("webservervue"),
	hydra.WithClusterName("taosytest"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.Web(":8072").Static(
		static.WithRoot("./src"),
		static.WithArchive("staticfile"),
		static.WithExts(".so", ".exe", ".pdf", ".txt", ".zip", ".gz", ".7z", ".tar", ".war", ".html", ".js", ".css", ".htm", ".ico", ".png", ".jpg", ".jpeg", ".md"),
		static.WithExclude(".png", ".exe", "/press/file7.gz", "/press/"),
		static.WithHomePage("index.html"),
		static.WithPrefix("/taosytest"),
		static.WithRewriters("/file5", "/"))

	hydra.Conf.Vars().Custom("config", "vue", map[string]interface{}{
		"api_addr":         fmt.Sprintf("//%s:50002", global.LocalIP()),
		"version":          time.Now().Format("20060102150405"),
		"currentComponent": "webvue",
	})
	app.Web("/vue/config", func(ctx hydra.IContext) interface{} {
		data := map[string]interface{}{}
		ctx.APPConf().GetVarConf().GetObject("config", "vue", &data)
		return data
	})
}

//webserver_vue 设置自定义static配置，使用vue和其他类型静态文件提供纯静态服务器demo

//1.1 运行程序 ./webservervue02 run

//1. http://192.168.5.94:8072/ 直接跳转道index.htm首页
//2. http://192.168.5.94:8072/file5 直接跳转道index.html首页
//3. http://192.168.5.94:8072/index.htm 直接跳转道index.html首页
//4. http://192.168.5.94:8072/index.html 直接返回index.html首页
//4. http://192.168.5.94:8072/default.html 直接返回default.html首页
//5. http://192.168.5.94:8072/default.htm 直接返回default.htm首页
//6. http://192.168.5.94:8072/file1.so 返回file1.so文件信息
//7. http://192.168.5.94:8072/file3.pdf 浏览器直接打开pdf文档
//8. http://192.168.5.94:8072/file4.txt 直接获取文件信息
//9. http://192.168.5.94:8072/views/file1.html 正常返回文件信息
//10. http://192.168.5.94:8072/views/file2.txt 正常返回文件信息
//11. http://192.168.5.94:8072/views/file3.html 正常返回文件信息
//12. http://192.168.5.94:8072/view/file1.html 正常返回文件信息
//13. http://192.168.5.94:8072/view/file2.txt 正常返回文件信息
//14. http://192.168.5.94:8072/view/file3.html 正常返回文件信息
//15. http://192.168.5.94:8072/web/file1.html 正常返回文件信息
//16. http://192.168.5.94:8072/web/file2.txt 正常返回文件信息
//17. http://192.168.5.94:8072/web/file3.html 正常返回文件信息
//18. http://192.168.5.94:8072/press/file6.zip 直接下载压缩文件信息
//19. http://192.168.5.94:8072/press/file7.gz 直接下载压缩文件信息
//20. http://192.168.5.94:8072/press/file8.7z 直接下载压缩文件信息
//21. http://192.168.5.94:8072/press/file9.tar.gz 直接下载压缩文件信息
//22. http://192.168.5.94:8072/press/file10.tar 直接下载压缩文件信息
//23. http://192.168.5.94:8072/press/file11.war 直接下载压缩文件信息
//24. http://192.168.5.94:8072/image/img2.jpg 正常返回图片
//25. http://192.168.5.94:8072/image/img3.jpg 正常返回图片
//26. http://192.168.5.94:8072/image/img4.jpeg 正常返回图片
//27. http://192.168.5.94:8072/taosytest/image/img4.jpeg 带前缀正常返回图片
//28. http://192.168.5.94:8072/taosytestimage/img4.jpeg 带前缀正常返回图片

//29. http://192.168.5.94:8072/press/ 状态码：404
//30. http://192.168.5.94:8072/view/ 状态码：404
//31. http://192.168.5.94:8072/views/ 状态码：404
//32. http://192.168.5.94:8072/web/ 状态码：404
//33. http://192.168.5.94:8072/file2.exe 状态码：404
//34. http://192.168.5.94:8072/press/file7.gz 状态码：404
//35. http://192.168.5.94:8072/image/img1.png 状态码：404

func main() {
	app.Start()
}
