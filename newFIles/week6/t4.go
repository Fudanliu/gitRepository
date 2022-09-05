package main

import "fmt"

//4. 自己实现一个BufferedFileWriter。
type BufferedFileWriter struct{
	buf [100]byte
	n int
}
//大部分情况下只是把cont写到缓冲区(go 变量)，当缓冲区写满之后才会触发真正的写磁盘操作
func (B *BufferedFileWriter) Write(cont string){
	bs := []byte(cont)
	l := len(bs)
	index := 0
	if l + B.n > 100{
		fmt.Printf("缓冲区已满，写入内存%s%s\n",B.buf,bs[0:100-B.n])
		l = l - (100 - B.n)
		index = 100 - B.n
		B.n = 0
	}
	
	for  i:= B.n; i< B.n + l; i++{
		B.buf[i] = bs[index]
		index++ 
	}
	B.n += l
	fmt.Printf("此次写入缓冲区数据为：%s, 共写入%d数据\n",B.buf[0:B.n],l)

}  

func main4(){
	B := new(BufferedFileWriter)

	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
	B.Write("12345678910")
}