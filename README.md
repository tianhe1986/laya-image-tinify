# laya-image-tinify
用GO语言编写，用于帮助laya小游戏项目处理png和jpg图片压缩。因为是远程调用[tinypng.com](https://tinypng.com/)的接口，因此暂时只支持这两种格式。

运行时需设置好四个目录（见下一节代码示例），源目录请设置为laya项目发布目录。

开始执行时，会先读取源目录下的`version.json`文件，得到所有要处理的文件，如果不存在此文件，则整个源目录都会被处理。

最终生成的目标目录结构会与源目录一致，其中非图片文件会直接复制，图片文件会进行压缩。

# 运行方式
1. 首先需要搭建golang环境，这个我就不详细讲了，随便搜一下吧。

2. 创建任一文件夹，并在其中创建main.go文件，示例如下：
```
package main

import (
	"github.com/tianhe1986/laya-image-tinify/tinifytool"
)

func main() {
	// api key
	var apiKey string = "pBnpqQTq35D1SPTQGJzgMhks620bf1Sp"

	// 要处理的目录，如果包含version文件，则只处理version文件中包含的文件，否则递归遍历所有文件
	var sourcePath string = "./srcDir"

	// 最终拷贝目录
	var finalPath string = "./finalDir"

	// 处理完的文件储存目录
	var handledPath string = "./test/handleDir"

	// 中间目录，防止直接处理源目录文件
	var tempPath string = "./test/tempDir"

	tinifytool.DoTinify(apiKey, sourcePath, finalPath, handledPath, tempPath)
}
```

其中`sourcePath`设置为laya项目发布目录， `finalPath`是最终处理完后生成的目录，`handledPath`和`tempPath`随便指定两个不同的空目录即可。

3. go get .

4. go run main.go
