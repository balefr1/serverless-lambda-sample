package service

import (
	"io"
	"log"
	filesprovider "sample_app/files_provider"
	"sample_app/model"
	"sample_app/repository"
	"time"
)

type FileService struct {
	filesRepository repository.FilesRepository
	filesProvider   filesprovider.FileProvider
}

func NewFileService(filesRepository repository.FilesRepository, filesProvider filesprovider.FileProvider) *FileService {
	return &FileService{filesRepository: filesRepository, filesProvider: filesProvider}
}

func (svc *FileService) FindFiles() ([]model.FileEntity, error) {
	return svc.filesRepository.FindAll()
}

func (svc *FileService) FindFile(name string) (*model.FileInfoDTO, error) {
	entity, err := svc.filesRepository.Find(name)
	if err != nil {
		return nil, err
	}
	info, err := svc.filesProvider.StatFile(name)
	if err != nil {
		return nil, err
	}
	return model.NewFileInfoDTO(entity, info.Size, info.Uid), nil
}

func (svc *FileService) SaveFile(name, ip string, file io.Reader) error {
	err := svc.filesRepository.Save(&model.FileEntity{FileName: name, FromIP: ip, CreatedAt: model.CustomTime(time.Now())})
	if err != nil {
		log.Println(err)
		return err
	}
	err = svc.filesProvider.UploadFile(name, file)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (svc *FileService) DownloadFile(name string) (string, error) {
	return svc.filesProvider.GetDownloadUrl(name)
}
