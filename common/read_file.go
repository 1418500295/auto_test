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
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	return strings.Trim(strings.SplitAfter(path, "auto_test")[0], " ")
}

func GetTestData(fileName string, caseIndex int) map[string]interface{} {
	defer func() {
		err3 := recover()
		if err3 != nil {
			fmt.Println(err3)
		}
	}()
	//byteData, err := ioutil.ReadFile(projectPath + "/testdata/" + fileName)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//file, err := f.Open(filepath.Join(filePath(), "/src/common/testdata/", fileName))
	file, err := f.Open("data/" + fileName)
	//file, err := os.Open(filepath.Join(FilePath(), "/data/"+fileName))
	if err != nil {
		fmt.Println(err)
	}
	//defer file.Close()

	reader := bufio.NewReader(file)
	var chunks []byte
	buf := make([]byte, 1024)
	var jsonData []map[string]interface{}
	for {
		n, err2 := reader.Read(buf)
		//io.EOF表示文件结束的错误
		if err2 != nil && err2 != io.EOF {
			panic(err2)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf...)
		err1 := json.Unmarshal(chunks[:n], &jsonData)
		if err1 != nil {
			fmt.Println(err1)
		}
	}
	return jsonData[caseIndex]
}

func GetApiUrl(urlName string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到异常: ", err)
		}
	}()
	//files, err1 := os.Open(filepath.Join(filePath(), "/src/common/application.properties"))
	//files, err1 := os.Open(filepath.Join(FilePath(), "/application.properties"))
	files, err1 := f.Open("application.properties")
	//defer files.Close()
	if err1 != nil {
		fmt.Println(err1)
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
