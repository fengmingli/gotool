package filesystem

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/**
 * @Author: LFM
 * @Date: 2022/7/24 19:51
 * @Since: 1.0.0
 * @Desc: TODO
 */

func AppendToFile(filePath string, content string) error {
	// 以只写的模式，打开文件
	fs, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file create failed. err - {%v} ", err)
	} else {
		// 查找文件末尾的偏移量
		n, _ := fs.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = fs.WriteAt([]byte(content), n)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(fs)
	return err
}

func CheckFileContainsStr(filePath, target string) bool {
	// 读取需要去重的文件内容
	f, _ := os.Open(filePath)
	defer func() {
		fsErr := f.Close()
		if fsErr != nil {
		}
	}()

	reader := bufio.NewReader(f)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if isPrefix {
			continue
		}
		if strings.Contains(string(line), target) {
			return true
		}
	}
	return false
}

//DistinctFile 为指定文件去重
func DistinctFile(file string, output string) error {
	// 读取需要去重的文件内容
	f, _ := os.Open(file)
	defer func() {
		fsErr := f.Close()
		if fsErr != nil {
		}
	}()
	reader := bufio.NewReader(f)
	// 去重map
	var set = make(map[string]bool, 1000)
	// 去重后的结果
	var result strings.Builder
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			lineStr := string(line)
			// key存在则跳出本次循环
			if set[lineStr] {
				continue
			}
			result.Write([]byte(fmt.Sprintf("%s\n", lineStr)))
			set[lineStr] = true
		}
	}
	// 写入另一个文件
	nf, _ := os.Create(output)
	_, err := io.Copy(nf, strings.NewReader(result.String()))
	if err != nil {
		return fmt.Errorf("file copy fail->{%v}", err)
	}
	defer func(nf *os.File) {
		fsErr := nf.Close()
		if fsErr != nil {
		}
	}(nf)
	return nil
}

// ExistsDirOrFile 判断所给路径文件/文件夹是否存在
func ExistsDirOrFile(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
