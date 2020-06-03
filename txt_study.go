package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main2 () {
	fmt.Println("-------")
	/**
	* 写txt文件
	 */
	/*file ,err := os.OpenFile("../s20200514/user.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	//及时关闭
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i:=0;i<3;i++{
		writer.WriteString("{\"name\":\"张三\",\"age\":1,\"sex\":\"男\"}\n")
	}
	writer.Flush()*/

	/**
	* 读取txt文件
	 */
	file, err := os.Open("../s20200514/user.txt")
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Println(str)
	}
	fmt.Println("读取结束了-----")
}
