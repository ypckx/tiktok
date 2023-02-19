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

	// 获取视频封面图片
	imageFile, err := GetImageFile(saveFile)

	videoFile := strings.Split(imageFile, ".")[0]
	videoFile = videoFile + ".mp4"
	if err != nil {
		return nil, err
	}

	// 向数据库中存放文件名路径，不包含ip地址（相对路径）
	err = model.InsertVideo(userId, videoFile, imageFile, title)
	if err != nil {
		return nil, err
	}

	return &common.PublishActionResponse{}, nil
}

func PublishList(tokenUserId, userId int64) (*common.PublishListResponse, error) {
	// 获取用户所有发布的视频
	videos, err := model.GetVideoList(userId)
	if err != nil {
		return nil, err
	}
	list := &common.PublishListResponse{
		VideoList: VideoList(videos, tokenUserId),
	}

	return list, nil
}

// 获取视频封面
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

	return videoName, nil
}
