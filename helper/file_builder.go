package helper

import (
	"bytes"
	"encoding/base64"
	"file-storage/app/exception"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)
func JpgBuilder(imgBase64 io.Reader) string {
	fileName := strconv.Itoa(int(time.Now().UnixNano()))+".jpeg"
	jpg,err := jpeg.Decode(imgBase64)
	if err != nil {
		fmt.Println(err.Error())
	}
	file,err := os.Create("./files/"+fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	jpeg.Encode(file,jpg,&jpeg.Options{
		Quality: 75,
	})
	return fileName
}

func PngBuilder(imgBase64 io.Reader) string {
	fileName := strconv.Itoa(int(time.Now().UnixNano()))+".jpeg"
	pngImg,err := png.Decode(imgBase64)
	if err != nil {
		fmt.Println(err.Error())
	}
	file,err := os.Create("./files/"+fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	png.Encode(file,pngImg)
	return fileName
}

func GetReader(dataUri string) io.Reader {
	base64Img := strings.Split(dataUri, ",")
	decoded,err := base64.StdEncoding.DecodeString(base64Img[1])
	exception.PanicIfBadRequest(err)
	reader := bytes.NewReader(decoded)
	return reader
}
