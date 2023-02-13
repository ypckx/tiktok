package service

import (
	"os/exec"
	"path/filepath"
	"strings"
	"tiktok/common"
	"tiktok/config"
	"tiktok/model"
)

func PublishVideo(userId int64, saveFile, title string) (*common.PublishActionResponse, error) {
	// client := minioStore.GetMinio()
	// videourl, err := client.UploadFile("video", saveFile, strconv.FormatInt(userId, 10))

	// if err != nil {
	// 	return nil, err
	// }
	imageFile, err := GetImageFile(saveFile)

	videoFile := strings.Split(imageFile, ".")[0]
	videoFile = videoFile + ".mp4"
	if err != nil {
		return nil, err
	}
	// fmt.Println("++++++[publishService.go PublishVideo] imageFile:", imageFile, "  videoFile:", videoFile)

	// return nil, errors.New("test image...")
	// // log.Debugf("imageFile %v\n", imageFile)
	// fmt.Printf("[func-PublishVideo] imageFile %v\n", imageFile)

	// picurl, err := client.UploadFile("pic", imageFile, strconv.FormatInt(userId, 10))
	// if err != nil {
	// 	picurl = "https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/7909abe413ec4a1e82032d2beb810157~tplv-k3u1fbpfcp-zoom-in-crop-mark:1304:0:0:0.awebp?"
	// }

	// 数据库中存放的为文件名，不包含url路径

	err = model.InsertVideo(userId, videoFile, imageFile, title)
	if err != nil {
		return nil, err
	}

	return &common.PublishActionResponse{}, nil
}

func PublishList(tokenUserId, userId int64) (*common.PublishListResponse, error) {
	// videos, err := model.GetVideoList(userId)  =====
	videos, err := model.GetVideoList(userId)
	if err != nil {
		return nil, err
	}
	list := &common.PublishListResponse{
		VideoList: VideoList(videos, tokenUserId),
	}

	return list, nil
}

func GetImageFile(videoPath string) (string, error) {
	temp := strings.Split(videoPath, "\\") // windows下使用\ ，linux下使用/
	videoName := temp[len(temp)-1]
	b := []byte(videoName)
	videoName = string(b[:len(b)-3]) + "jpg"
	picpath := config.PicPath
	picName := filepath.Join(picpath, videoName)

	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", picName)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	// log.Debugf(picName)
	// fmt.Println("[publishService func-GetImageFile] ", picName)
	return videoName, nil
}
