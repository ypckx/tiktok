package service

import (
	"tiktok/common"
	"tiktok/model"
)

func CommentAction(commentId, videoId, userId int64, comment_text, actionType string) (*common.CommentActionResponse, error) {

	if actionType == "1" {
		//commentInfo, err := repository.CommentAdd(userId, videoId, comment_text)
		commentInfo, err := model.CommentAdd(userId, videoId, comment_text)

		if err != nil {
			return nil, err
		}
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
	comments, err := model.CommentList(videoId)
	if err != nil {
		return nil, err
	}
	// log.Debugf("comments:%v\n", comments)
	// fmt.Printf("[func-CommentList] comments:%v\n", comments)

	list := &common.CommentListResponse{
		CommentList: make([]*common.Comment, len(comments)),
	}

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
