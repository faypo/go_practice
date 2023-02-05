package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var x string
	var y string

	path := "."
	files, _ := ioutil.ReadDir(path)

	fmt.Println("請輸入檔名更需更改的字")
	fmt.Scanln(&x)
	fmt.Println("更改為")
	fmt.Scanln(&y)

	for index, f := range files {
		// 帶副檔名的檔名
		fullFilename := f.Name()
		fmt.Println(index)
		fmt.Println(fullFilename)
		//副檔名
		fileExt := filepath.Ext(f.Name())
		fmt.Println(fileExt)
		// 不帶副檔名的檔名
		filenameOnly := strings.TrimSuffix(fullFilename, fileExt)
		fmt.Println(filenameOnly)
		//將每個檔名後面加上1，副檔名不變
		//os.Rename(path+"\\"+f.Name(), path+"\\"+fmt.Sprintf("%s%s%s", filenameOnly, "1", fileExt))
		//將每個檔名中的1替換為2，副檔名不變
		os.Rename(path+"\\"+f.Name(), path+"\\"+fmt.Sprintf("%s%s", strings.Replace(filenameOnly, x, y, -1), fileExt))
	}

	//new = strings.Replace(x, y, "8", -1)
	//fmt.Printf("哈哈: %s", new)
	fmt.Println("搞定拉美眉")
	fmt.Println("請按enter鍵退出....")
	fmt.Scanln(&x)
	fmt.Println("")
}
