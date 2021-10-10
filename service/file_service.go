package service

type FileService interface {
	Save(files []string) []string
}