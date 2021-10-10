package service

import (
	"file-storage/app"
	"file-storage/helper"
	"strings"
)

type fileServiceImpl struct{}

func NewFileServiceImpl() FileService {
	return fileServiceImpl{}
}

func (f fileServiceImpl) Save(files []string) []string {
	result := []string{}

	for _, v := range files {
		items := strings.Split(v,";")
		if items[0] == "data:image/jpg" || items[0] == "data:image/jpeg" {
			reader := helper.GetReader(items[1])
			fileName := helper.JpgBuilder(reader)
			result = append(result,app.HOST+fileName)
		} else if items[0] == "data:image/png" {
			reader := helper.GetReader(items[1])
			fileName := helper.PngBuilder(reader)
			result = append(result,app.HOST+fileName)
		}
	}
	return result
}

