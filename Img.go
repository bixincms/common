package common

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//将图片base数据装换成图片
//Base64 to ImgFile
func ImgBase64ToImg(inputBase64Str string, saveFileName string) (e error) {
	img_split := regexp.MustCompile(`image/.*base64,(.*)`).FindStringSubmatch(inputBase64Str)
	if len(img_split) != 2 {
		e = errors.New("the no image/*")
		return
	}
	img_byte, _ := base64.StdEncoding.DecodeString(img_split[1])
	e = ioutil.WriteFile(saveFileName, img_byte, 0666)
	return
}

//图片转换成base64
//Convert pictures to Base64
func ImgToBase64(SourceFileName string) (e error, s string) {
	file, e := os.Open(SourceFileName)
	if e != nil {
		return
	}
	defer file.Close()
	contentType, e := GetFileContentType(file)
	if e != nil {
		e = errors.New("the no image/*")
		return
	}
	//判断mimetype是否是Images
	contentType = strings.ToLower(contentType)
	if !strings.Contains(contentType, "image/") {
		return
	}
	ctype := strings.ReplaceAll(contentType, "image/", "")
	r2, _ := ioutil.ReadFile(SourceFileName)
	toString := base64.StdEncoding.EncodeToString(r2)
	s = fmt.Sprintf("data:image/%s;base64,%s", ctype, toString)
	return
}
func GetFileContentType(out *os.File) (string, error) {

	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
