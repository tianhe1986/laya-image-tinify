package main

import (
	"laya-image-tinify/tinifytool"
)

func main() {
	// api key
	var apiKey string = ""

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
