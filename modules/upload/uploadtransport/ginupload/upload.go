package ginupload

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/upload/uploadbiz"
	"bestHabit/modules/upload/uploadstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(imgStore, appCtx.UploadProvider())

		user := c.MustGet(common.CurrentUser).(common.Requester)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename, user.GetId())

		if err != nil {
			panic(err)
		}

		//upload in root directory (server)
		//c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
