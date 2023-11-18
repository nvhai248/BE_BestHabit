package uploadbiz

import (
	"bestHabit/component/uploadprovider"
	"bestHabit/modules/upload/uploadmodel"
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type CreateStorage interface {
	Create(ctx context.Context, data *uploadmodel.ImageUpload) error
}

type uploadBiz struct {
	store      CreateStorage
	upProvider uploadprovider.UploadProvider
}

func NewUploadBiz(store CreateStorage, upProvider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{store: store, upProvider: upProvider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string, userId int) (*uploadmodel.ImageUpload, error) {
	/* reader := bytes.NewReader(data)

	imgConfig, _, err := image.Decode(reader)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}
	imgConfig.Bounds().Dx() */

	w := 0
	h := 0

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" -> "jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9128213314.jpg

	img, err := biz.upProvider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	var dataSave uploadmodel.ImageUpload

	dataSave.CloudName = img.CloudName
	dataSave.Extension = img.Extension
	dataSave.Width = img.Width
	dataSave.Height = img.Height
	dataSave.Url = img.Url
	dataSave.CreatedBy = userId

	if err := biz.store.Create(ctx, &dataSave); err != nil {
		// delete image S3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return &dataSave, nil
}
