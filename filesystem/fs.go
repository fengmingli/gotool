package filesystem

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

// RemoveFile 删除指定的文件
func RemoveFile(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if info != nil && !info.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", filePath)
	}

	if err = os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to remove file: %w", err)
	}
	return nil
}

// RemoveFilesInDir 删除指定目录及其所有内容
func RemoveFilesInDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil {
		return fmt.Errorf("failed to get directory info: %w", err)
	}

	if info != nil && !info.IsDir() {
		return fmt.Errorf("%s is not a directory", dirPath)
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to list files in directory: %w", err)
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		if file.IsDir() {
			err := RemoveFilesInDir(filePath)
			if err != nil {
				return fmt.Errorf("failed to remove directory %s: %w", filePath, err)
			}
		} else {
			err := os.Remove(filePath)
			if err != nil {
				return fmt.Errorf("failed to remove file %s: %w", filePath, err)
			}
		}
	}
	return nil
}

// DeleteFolder 用于删除指定的文件夹及其内容
func DeleteFolder(folderPath string, needDeleteCurrentFolder bool) error {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil && info != nil {
			return err
		}
		// 跳过最外层的文件夹
		if !needDeleteCurrentFolder && path == folderPath {
			return nil
		}

		// 检查路径是否是直接子文件或子文件夹
		relativePath, err := filepath.Rel(folderPath, path)
		if err != nil {
			fmt.Println("Error getting relative path:", path, ":", err)
			return err
		}
		components := strings.Split(relativePath, string(filepath.Separator))
		if len(components) == 1 {
			// 如果是文件夹，使用os.RemoveAll删除文件夹及其内容
			if info.IsDir() {
				err = os.RemoveAll(path)
				if err != nil {
					return err
				}
			} else {
				// 如果是文件，直接删除
				err = os.Remove(path)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	// 检查删除过程中是否有错误发生
	if err != nil {
		return err
	}

	return nil
}
