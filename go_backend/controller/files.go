package controller

import (
	"log"
	"net/http"
	"sample_app/service"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	filesService *service.FileService
}

func NewFileController(filesService *service.FileService) *FileController {
	return &FileController{filesService: filesService}
}

func (controller *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		handleError(err, "error reading file", c)
		log.Println(err)
		return
	}
	multipFile, err := file.Open()
	if err != nil {
		handleError(err, "error reading file", c)
		return
	}
	err = controller.filesService.SaveFile(file.Filename, c.GetHeader("X-Forwarded-For"), multipFile)
	if err != nil {
		handleError(err, "", c)
		return
	}
	log.Println("uplaoded file " + file.Filename)
	c.JSON(201, gin.H{"name": file.Filename, "status": "ok"})
}

func (controller *FileController) FindFiles(c *gin.Context) {

	res, err := controller.filesService.FindFiles()
	if err != nil {
		log.Println(err)
		handleError(err, "", c)
		return
	}
	c.JSON(200, res)
}

func (controller *FileController) DownloadFile(c *gin.Context) {
	fileName := c.Param("file")
	provFile, err := controller.filesService.FindFile(fileName)
	if err != nil {
		log.Println(err)
		handleError(err, "", c)
		return
	}
	res, err := controller.filesService.DownloadFile(fileName)
	if err != nil {
		log.Println(err)
		handleError(err, "", c)
		return
	}
	c.JSON(200, gin.H{"name": fileName, "download_url": res, "size": provFile.Size, "etag": provFile.Etag})
}

func BodySizeMiddleware(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, 10*1024*1024)
	c.Next()
}

func handleError(err error, externalMessage string, c *gin.Context) {
	if externalMessage == "" {
		externalMessage = "generic server error"
	}
	c.AbortWithStatusJSON(500, gin.H{"error_message": externalMessage})
}
