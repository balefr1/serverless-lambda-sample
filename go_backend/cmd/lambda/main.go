package main

import (
	"context"
	"os"
	"sample_app/controller"
	filesprovider "sample_app/files_provider"
	"sample_app/repository"
	"sample_app/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	session := session.New()
	filesRepo := repository.NewDynamoFilesRepository(session, os.Getenv("DYNAMO_TABLE"))
	filesProvider := filesprovider.NewS3FileProvider(session, os.Getenv("S3_DIR"), os.Getenv("S3_BUCKET"))
	filesService := service.NewFileService(filesRepo, filesProvider)
	filesController := controller.NewFileController(filesService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(controller.BodySizeMiddleware)
	router.Use(cors.Default())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"ping": "ok"})
	})
	router.GET("/files", filesController.FindFiles)
	router.POST("/files", filesController.UploadFile)
	router.POST("/file/:file/download", filesController.DownloadFile)

	ginLambda := ginadapter.New(router)
	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return ginLambda.ProxyWithContext(ctx, req)
	})
}
