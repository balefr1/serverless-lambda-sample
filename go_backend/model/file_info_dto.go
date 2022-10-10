package model

type FileInfoDTO struct {
	FileEntity
	Size int64  `json:"size"`
	Etag string `json:"etag"`
}

func NewFileInfoDTO(fileEntity *FileEntity, size int64, etag string) *FileInfoDTO {
	return &FileInfoDTO{FileEntity: *fileEntity, Size: size, Etag: etag}
}
