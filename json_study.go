package main

import(
	"fmt"
	//"encoding/json"
	//"os"
)

type Website struct {
	Name string
	Url string
	Demo []string
}

func main1()  {
	fmt.Println("aaaa")
	/**
	* 生成JSON文件
	*/
	/*info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

    // 创建文件
    filePtr, err := os.Create("info.json")
    if err != nil {
        fmt.Println("文件创建失败", err.Error())
        return
    }
    defer filePtr.Close()

    // 创建Json编码器
    encoder := json.NewEncoder(filePtr)

    err = encoder.Encode(info)
    if err != nil {
        fmt.Println("编码错误", err.Error())

    } else {
        fmt.Println("编码成功")
	}*/
	
	/**
	* 读取JSON文件
	*/
	/*filePtr,err := os.Open("info.json")
	if err == nil {
		defer filePtr.Close()
		var info []Website
		// 创建json解码器
		decoder := json.NewDecoder(filePtr)
		err = decoder.Decode(&info)
		if err != nil {
			fmt.Println("解码失败", err.Error())
		} else {
			fmt.Println("解码成功")
			fmt.Println(info)
		}
	}*/
}