package handler

import (
	"atro/internal/helper"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//FileUploadHandler -> Interface to File Upload
type FileUploadHandler interface {
	SingleFile(*gin.Context)
}

//SingleFile --> handle uploading of single file
func SingleFile(c *gin.Context) {

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "gửi lên tao có thấy gì đâu ?", err.Error()))
		return
	}

	savePath := os.Getenv("IMAGE_SAVE_PATH")
	log.Println(file.Filename+"save at: "+ savePath)

	err = c.SaveUploadedFile(file, savePath+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "lấy dc ảnh nhưng lỗi cmnr khi savefile", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(1, "get file done", "file-name:"+file.Filename))

}
