package models

/*******************************
* 响应数据类型
*******************************/
// Golang 的结构体定义中添加 omitempty 关键字，
// 来表示这条信息如果没有提供，在序列化成 json 的时候就不要包含其默认值

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

/*
{
    "status_code": 0,
    "status_msg": "string",
    "next_time": 0,
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
*/
type FeedResp struct {
	Response
	NextTime  int64   `json:"next_time,omitempty"`
	VideoList []Video `json:"video_list,omitempty"`
}

/*
	{
	    "status_code": 0,
	    "status_msg": "string",
	    "user_id": 0,
	    "token": "string"
	}
*/
type RegisterResp struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

/*
	{
	    "status_code": 0,
	    "status_msg": "string",
	    "user_id": 0,
	    "token": "string"
	}
*/
type LoginResp struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

/*
	{
	    "status_code": 0,
	    "status_msg": "string",
	    "user": {
	        "id": 0,
	        "name": "string",
	        "follow_count": 0,
	        "follower_count": 0,
	        "is_follow": true
	    }
	}
*/
type UserInfoResp struct {
	Response
	UserInfo User `json:"user"`
}

/*
	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
type PublishActionResp struct {
	Response
}

/*
{
    "status_code": 0,
    "status_msg": "string",
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
*/

type PublishListResp struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

/*
	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
type FavoriteActionResp struct {
	Response
}

/*
	{
	    "status_code": "string",
	    "status_msg": "string",
	    "video_list": [
	        {
	            "id": 0,
	            "author": {
	                "id": 0,
	                "name": "string",
	                "follow_count": 0,
	                "follower_count": 0,
	                "is_follow": true
	            },
	            "play_url": "string",
	            "cover_url": "string",
	            "favorite_count": 0,
	            "comment_count": 0,
	            "is_favorite": true,
	            "title": "string"
	        }
	    ]
	}
*/
type FavoriteListResp struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}

/*
{
    "status_code": 0,
    "status_msg": "string",
    "comment": {
        "id": 0,
        "user": {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true
        },
        "content": "string",
        "create_date": "string"
    }
}
*/

type CommentActionResp struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

/*
{
    "status_code": 0,
    "status_msg": "string",
    "comment_list": [
        {
            "id": 0,
            "user": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true
            },
            "content": "string",
            "create_date": "string"
        }
    ]
}
*/

type CommentListResp struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

/*
	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
type RelationActionResp struct {
	Response
}

/*
	{
	    "status_code": "string",
	    "status_msg": "string",
	    "user_list": [
	        {
	            "id": 0,
	            "name": "string",
	            "follow_count": 0,
	            "follower_count": 0,
	            "is_follow": true
	        }
	    ]
	}
*/
type FollowListResp struct {
	Response
	UserList []User `json:"user_list,omitempty"`
}

/*
	{
	    "status_code": "string",
	    "status_msg": "string",
	    "user_list": [
	        {
	            "id": 0,
	            "name": "string",
	            "follow_count": 0,
	            "follower_count": 0,
	            "is_follow": true
	        }
	    ]
	}
*/
type FollowerListResp struct {
	Response
	UserList []User `json:"user_list,omitempty"`
}

/*
	{
	    "status_code": "string",
	    "status_msg": "string",
	    "user_list": [
	        {
	            "id": 0,
	            "name": "string",
	            "follow_count": 0,
	            "follower_count": 0,
	            "is_follow": true
	        }
	    ]
	}
*/
type FriendListResp struct {
	Response
	UserList []User `json:"user_list,omitempty"`
}

/*
	{
	    "status_code": 0,
	    "status_msg": "string"
	}
*/
type MessageActionResp struct {
	Response
}

/*
	{
	    "status_code": "string",
	    "status_msg": "string",
	    "message_list": [
	        {
	            "id": 0,
	            "content": "string",
	            "create_time": "string"
	        }
	    ]
	}
*/
type MessageChatResp struct {
	Response
	MessageList []Message `json:"message_list,omitempty"`
}
