package filesystem

import (
	"fmt"
	"testing"
)

/**
 * @Author: LFM
 * @Date: 2022/7/24 19:57
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestAppendToFile(t *testing.T) {
	err := AppendToFile("/tmp/fstab", "/dev/mapper/xxx /mnt/lvm-simple/xxxxss defaults 0 \r\n")
	if err != nil {
		fmt.Println(err)
	}
}

func TestDistinctFile(t *testing.T) {
	err := DistinctFile("/tmp/fstab", "/tmp/fstab-1")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsDir(t *testing.T) {
	fmt.Println(IsDir("/tmp/fstab"))
}

func TestIsFile(t *testing.T) {
	fmt.Println(IsFile("/tmp/fstab"))
}

func TestExistsDirOrFile(t *testing.T) {
	fmt.Println(ExistsDirOrFile("/tmp/fstab"))
}

func TestCheckFileContainsStr(t *testing.T) {
	fmt.Println(CheckFileContainsStr("/tmp/fstab", "/dev/mapper/xxxss"))
}
