package main

import (
	"flag"
	"fmt"
	"github.com/kutear/Fuck-Hard-Code/utils"
	_ "github.com/kutear/Fuck-Hard-Code/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	inPath := flag.String("input", "", "The Path Of Layout Root")
	outPath := flag.String("output", "", "文件输出目录")
	//prefix := flag.String("p", "", "指定输出到strings.xml的key值前缀")
	flag.Parse()
	if *inPath == "" || *outPath == "" {
		fmt.Println("please input -h to see usage")
		return
	}

	utils.CreatePath(*outPath)
	err := filepath.Walk(*inPath, func(file string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println("跳过目录" + file)
			return nil
		}
		utils.DealFile(file, info.Name(), *outPath)
		return nil
	})

	outResPath := *outPath + string(os.PathSeparator) + "out"
	utils.CreatePath(outResPath)
	str := "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n\n"
	ioutil.WriteFile(outResPath+string(os.PathSeparator)+"dimens.xml", []byte(str+utils.GetDimen().ChildSingleLineOut()), 0644)
	ioutil.WriteFile(outResPath+string(os.PathSeparator)+"strings.xml", []byte(str+utils.GetString().ChildSingleLineOut()), 0644)

	if err != nil {
		fmt.Println(err.Error())
	}

}
