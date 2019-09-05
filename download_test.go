package common

import (
	"fmt"
	"testing"
)

func TestDownload(t *testing.T) {
	n, err := Download("http://img.365jia.cn/uploads/16/0614/.575fb6acb0619t2048l90.jpg", "static/575fb6acb0619t2048l90.jpg")
	fmt.Println(n, err)
}
