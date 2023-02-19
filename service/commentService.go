package service

import (
	"tiktok/common"
	"tiktok/model"
)

func CommentAction(commentId, videoId, userId int64, comment_text, actionType string) (*common.CommentActionResponse, error) {

	// actionType = 1。1-发布评论，2-删除评论
	if actionType == "1" {
		// 添加评论
		commentInfo, err := model.CommentAdd(userId, videoId, comment_text)
		if err != nil {
			return nil, err
		}

		// 封装响应数据
		user, _ := model.GetUserInfo(userId)
		commentResponse := &common.CommentActionResponse{
			Comment: &common.Comment{
				Id:         commentInfo.CommentId,
				User:       messageUserInfo(user),
				Content:    comment_text,
				CreateDate: commentInfo.Time,
			},
		}

		return commentResponse, nil
	} else {
		err := model.CommentDelete(videoId, commentId)
		if err != nil {
			return nil, err
		}
		commentResponse := &common.CommentActionResponse{
			Comment: nil,
		}
		return commentResponse, nil
	}

}

// 用户评论
func CommentList(videoId int64) (*common.CommentListResponse, error) {
	// 获取评论列表
	comments, err := model.CommentList(videoId)
	if err != nil {
		return nil, err
	}

	list := &common.CommentListResponse{
		CommentList: make([]*common.Comment, len(comments)),
	}

	// 封装响应数据
	for i, comment := range comments {
		//为了找到video_id所对应的user_id，在通过user_id找到user_name.传递给前端
		userID := comment.UserId
		user, _ := model.GetUserInfo(userID)
		userinfo := messageUserInfo(user)

		v := &common.Comment{
			Id:         comment.CommentId,
			User:       userinfo,
			Content:    comment.Comment,
			CreateDate: comment.Time,
		}
		list.CommentList[i] = v
	}

	return list, nil
}
