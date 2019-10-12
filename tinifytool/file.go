package tinifytool

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// 复制文件，从srcPath复制到dstPath
func CopyFile(dstPath string, srcPath string) error {
	var err error
	// 文件句柄
	var srcfd *os.File
	var dstfd *os.File

	//打开要读的文件
	if srcfd, err = os.Open(srcPath); err != nil {
		return nil
	}
	defer srcfd.Close()

	//创建要写的文件
	if dstfd, err = os.Create(dstPath); err != nil {
		return nil
	}
	defer dstfd.Close()

	// 拷贝文件
	_, err = io.Copy(dstfd, srcfd)

	return err
}

// 递归复制文件夹，从srcPath复制到dstPath
func CopyDir(dstDir string, srcDir string) error {
	var err error
	var infos []os.FileInfo

	// 读源文件夹信息
	if infos, err = ioutil.ReadDir(srcDir); err != nil {
		return err
	}

	// 创建要写的文件夹
	if err = CreateDir(dstDir); err != nil {
		return err
	}

	// 依次遍历，递归处理复制
	for _, info := range infos {
		srcTempPath := path.Join(srcDir, info.Name())
		dstTempPath := path.Join(dstDir, info.Name())
		if info.IsDir() { //是目录，继续递归处理
			if err = CopyDir(dstTempPath, srcTempPath); err != nil {
				fmt.Println(err)
			}
		} else { //是文件，直接复制
			if err = CopyFile(dstTempPath, srcTempPath); err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}

// 创建文件夹
func CreateDir(path string) error {
	return os.MkdirAll(path, 0666)
}

// 递归创建同样结构的文件夹，只创建文件夹，不复制文件
func CreateSameDir(dstDir string, srcDir string) error {
	var err error
	var infos []os.FileInfo

	// 读源文件夹信息
	if infos, err = ioutil.ReadDir(srcDir); err != nil {
		return err
	}

	// 创建要写的文件夹
	if err = CreateDir(dstDir); err != nil {
		return err
	}

	// 依次遍历，递归处理复制
	for _, info := range infos {
		srcTempPath := path.Join(srcDir, info.Name())
		dstTempPath := path.Join(dstDir, info.Name())
		if info.IsDir() { //是目录，继续递归处理
			if err = CreateSameDir(dstTempPath, srcTempPath); err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}

// 文件是否存在
func IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
