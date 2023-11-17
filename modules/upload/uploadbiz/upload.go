package uploadbiz

import (
	"bestHabit/common"
	"bestHabit/component/uploadprovider"
	"bestHabit/modules/upload/uploadmodel"
	"bytes"
	"context"
	"fmt"
	"image"
	"path/filepath"
	"strings"
	"time"
)

type CreateStorage interface {
	Create(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	store      CreateStorage
	upProvider uploadprovider.UploadProvider
}

func NewUploadBiz(store CreateStorage, upProvider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{store: store, upProvider: upProvider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	reader := bytes.NewReader(data)

	imgConfig, _, err := image.DecodeConfig(reader)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	w := imgConfig.Width
	h := imgConfig.Height

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
	img.CloudName = "s3" //should be set in provider
	img.Extension = fileExt

	if err := biz.store.Create(ctx, img); err != nil {
		// delete image S3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil
}
