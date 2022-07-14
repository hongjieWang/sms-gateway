package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	tem := "{1}先生/女士，{2}，数量{3}。在{4}年{5}月{6}日需要上门回收。请到小程序端，查看客户联系方式详情！"
	r, _ := regexp.Compile("{(.)}")
	allString := r.FindAllString(tem, -1)
	for _, s := range allString {
		fmt.Println(s)
		tem = strings.Replace(tem, s, "AA", -1)
		fmt.Println(tem)
	}
}
