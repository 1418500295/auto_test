package common

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	//go:embed application.properties
	//go:embed data/*
	f embed.FS
)

func FilePath() string {
	path, err5 := os.Getwd()
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(path)
	return strings.Trim(strings.SplitAfter(path, "auto_test")[0], " ")
}

func GetTestData(fileName string, caseIndex int) map[string]interface{} {
	defer func() {
		err53 := recover()
		if err53 != nil {
			fmt.Println(err53)
		}
	}()
	//byteData, err5 := ioutil.ReadFile(projectPath + "/testdata/" + fileName)
	//if err5 != nil {
	//	fmt.Println(err5)
	//}
	//file, err5 := f.Open(filepath.Join(filePath(), "/src/common/testdata/", fileName))
	file, err5 := f.Open("data/" + fileName)
	//file, err5 := os.Open(filepath.Join(FilePath(), "/data/"+fileName))
	if err5 != nil {
		fmt.Println(err5)
	}
	//defer file.Close()

	reader := bufio.NewReader(file)
	var chunks []byte
	buf := make([]byte, 1024)
	var jsonData []map[string]interface{}
	for {
		n, err52 := reader.Read(buf)
		//io.EOF表示文件结束的错误
		if err52 != nil && err52 != io.EOF {
			panic(err52)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf...)
		err5 := json.Unmarshal(chunks[:n], &jsonData)
		if err5 != nil {
			fmt.Println(err5)
		}
	}
	return jsonData[caseIndex]
}

func GetApiUrl(urlName string) string {
	defer func() {
		err5 := recover()
		if err5 != nil {
			fmt.Println("捕获到异常: ", err5)
		}
	}()
	//files, err5 := os.Open(filepath.Join(filePath(), "/src/common/application.properties"))
	//files, err5 := os.Open(filepath.Join(FilePath(), "/application.properties"))
	files, err5 := f.Open("application.properties")
	//defer files.Close()
	if err5 != nil {
		fmt.Println(err5)
	}
	bytesStr, _ := io.ReadAll(files)
	configStr := strings.Split(string(bytesStr), "\n")
	var endUrl string
	var host string
	for _, i := range configStr {
		iSlice := strings.Split(strings.ReplaceAll(i, "\r", ""), "=")
		if "host" == strings.Trim(iSlice[0], " ") {
			host = strings.Trim(iSlice[1], " ")
		}
		if urlName == strings.Trim(iSlice[0], " ") {
			url := strings.Trim(iSlice[1], " ")
			if !strings.HasPrefix(url, "/") {
				endUrl = host + "/" + fmt.Sprintf("%v", url)
			} else if strings.HasPrefix(url, "/") {
				endUrl = host + fmt.Sprintf("%v", url)
			} else {
				fmt.Println("请求地址格式不正确")
			}
		}
	}
	return endUrl
}
