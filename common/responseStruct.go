package common

// var GetUserIdByToken = map[string]int64{
// 	"ypckx123":       1,
// 	"zhangleidouyin": 1,
// }

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        *User  `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       *User  `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackGroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  int64  `json:"total_favorited,omitempty"`
	WorkCount       int64  `json:"work_count,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
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

/***********************************/
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User *User `json:"user"`
}

type UserListResponse struct {
	Response
	UserList []*User `json:"user_list"`
}

type FeedResponse struct {
	Response
	VideoList []*Video `json:"video_list,omitempty"`
	NextTime  int64    `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
}

type CommentListResponse struct {
	Response
	CommentList []*Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment *Comment `json:"comment,omitempty"`
}

type FavoriteListResponse struct {
	Response
	VideoList []*Video `json:"video_list,omitempty"`
}

type FollowListResponse struct {
	Response
	UserList []*User `json:"user_list,omitempty"`
}

type FollowerListResponse struct {
	Response
	UserList []*User `json:"user_list,omitempty"`
}

type FriendListResponse struct {
	Response
	UserList []*User `json:"user_list,omitempty"`
}

type MessageActionResponse struct {
	Response
}

type MessageListResponse struct {
	Response
	MessageList []*Message `json:"message_list"`
}

type PublishActionResponse struct {
	Response
}

type PublishListResponse struct {
	Response
	VideoList []*Video `json:"video_list,omitempty"`
}
