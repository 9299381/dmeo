package readers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type IniReader struct {
}

func (s *IniReader) New() *IniReader {
	//这里可以做一些初始化
	return s
}

func (s *IniReader) Read(filePath string) interface{} {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	data := make(map[string]map[string]string)
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				s.CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#": //增加配置文件备注
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i+1:])
			data[section][strings.TrimSpace(line[0:i])] = value
		}
	}

	return data
}

func (s *IniReader) CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}
