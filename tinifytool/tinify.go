package tinifytool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/tianhe1986/tinify-go/tinify"
)

// 要处理的目录，如果包含version文件，则只处理version文件中包含的文件，否则递归遍历所有文件
var sourcePath string = ""

// 最终拷贝目录
var finalPath string = ""

// 处理完的文件储存目录
var handledPath string = ""

// 中间目录，防止直接处理源目录文件
var tempPath string = ""

var fileNameMap map[string]string

// 复制到临时目录
func copyToTemp() {
	CopyDir(tempPath, sourcePath)
}

func DoTinify(key string, srcPath string, finPath string, tmpHandlePath string, tmpPath string) {
	// 设置配置参数
	tinify.SetKey(key)
	sourcePath = srcPath
	finalPath = finPath
	handledPath = tmpHandlePath
	tempPath = tmpPath

	// 复制到临时目录
	copyToTemp()

	// 如果存在 version文件， 读入到fileNameMap，只处理fileNameMap中的文件
	var versionFilePath string = path.Join(tempPath, "version.json")
	if IsExist(versionFilePath) {
		byteValue, err := ioutil.ReadFile(versionFilePath)
		if err == nil {
			json.Unmarshal(byteValue, &fileNameMap)
			fileNameMap["version.json"] = "version.json"
		}
	}

	//建立相应的文件夹
	createSameHandleDir()

	if fileNameMap != nil { //只处理fileNameMap中的
		compressWithMap()
	} else { // 全部遍历
		compressFullDir()
	}

	// 复制到最终文件夹
	copyToFinal()
}

// 建立相应的文件夹
func createSameHandleDir() {
	CreateSameDir(handledPath, tempPath)
}

// 整个文件夹都遍历压缩
func compressFullDir() {
	compressTo(handledPath, tempPath)
}

// 根据映射文件遍历压缩图片
func compressWithMap() {
	for _, v := range fileNameMap {
		if strings.HasSuffix(v, ".png") || strings.HasSuffix(v, ".jpg") || strings.HasSuffix(v, ".jpeg") { // 是图片类型，调用远程压缩处理
			tinifyToHandle(v)
		} else { // 不是图片类型，直接复制
			copyToHandle(v)
		}
	}
}

// 递归遍历处理文件
func compressTo(dstDir string, srcDir string) error {
	var err error
	var infos []os.FileInfo

	// 读源文件夹信息
	if infos, err = ioutil.ReadDir(srcDir); err != nil {
		return err
	}

	// 依次遍历，递归处理复制
	for _, info := range infos {
		srcTempPath := path.Join(srcDir, info.Name())
		dstTempPath := path.Join(dstDir, info.Name())
		if info.IsDir() { //是目录，继续递归处理
			if err = compressTo(dstTempPath, srcTempPath); err != nil {
				fmt.Println(err)
			}
		} else { //是文件
			// 是图片类型，调用远程压缩处理
			if strings.HasSuffix(info.Name(), ".png") || strings.HasSuffix(info.Name(), ".jpg") || strings.HasSuffix(info.Name(), ".jpeg") {
				tinifyTo(dstTempPath, srcTempPath)
			} else {
				copyTo(dstTempPath, srcTempPath)
			}
		}
	}

	return nil
}

// 直接复制到处理后文件夹
func copyToHandle(filename string) {
	srcPath := path.Join(tempPath, filename)
	dstPath := path.Join(handledPath, filename)
	copyTo(dstPath, srcPath)
}

func copyTo(dstPath string, srcPath string) {
	// 如果已经存在，不再处理
	if !IsExist(dstPath) {
		CopyFile(dstPath, srcPath)
	}
}

func tinifyToHandle(filename string) {
	srcPath := path.Join(tempPath, filename)
	dstPath := path.Join(handledPath, filename)
	tinifyTo(dstPath, srcPath)
}

func tinifyTo(dstPath string, srcPath string) {
	// 如果已经存在，不再次请求
	if !IsExist(dstPath) {
		tinify.FromFile(srcPath).ToFile(dstPath)
	}

	fmt.Printf("tinify to %s.\n", dstPath)
}

// 复制到最终文件夹
func copyToFinal() {
	CopyDir(finalPath, handledPath)
}
