package util

import (
	"io/ioutil"
	"fmt"
	"os"
)

func Write2file(fileName, contents string) {
	data := []byte(contents)
	if ioutil.WriteFile(fileName, data, 0644) == nil {
		fmt.Println("写入文件成功", contents)
	}

}

func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
