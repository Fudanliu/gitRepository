package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//3. 把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩
func walk(path string) error {
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, fileInfo := range fileInfos {
			if strings.HasSuffix(fileInfo.Name(), ".txt") {
				deTxt(fileInfo.Name())
			} else if fileInfo.IsDir() {
				if err := walk(filepath.Join(path, fileInfo.Name())); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func deTxt(fileName string) {
	fin, _ := os.Open(fileName)
	reader := bufio.NewReader(fin)
	fout, _ := os.OpenFile("compress.zlib", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	writer := zlib.NewWriter(fout)
	for {
		if line, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					writer.Write([]byte(line))
				}
				break
			}
		} else {
			writer.Write([]byte(line))
		}
		writer.Flush()
	}
}
func main() {
	err := walk("./")
	fmt.Print(err)
}
