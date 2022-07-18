package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	tem := "${name}先生/女士，${age}，数量${phone}。在${4}年{5}月{6}日需要上门回收。请到小程序端，查看客户联系方式详情！"
	params := []string{"张三", "20", "18232533068", "2022"}
	r, _ := regexp.Compile("\\$\\{\\w+\\}")
	allString := r.FindAllString(tem, -1)
	var mapJson = make(map[string]string)
	for i, s := range allString {
		fmt.Println(s)
		r1, _ := regexp.Compile("[a-z]")
		findAllString := r1.FindAllString(s, -1)
		key := ""
		for _, s := range findAllString {
			key = key + s
		}
		mapJson[key] = params[i]
		fmt.Println(key)
		tem = strings.Replace(tem, s, "AA", -1)
		fmt.Println(tem)
	}
	marshal, _ := json.Marshal(mapJson)
	fmt.Println(string(marshal))
}
