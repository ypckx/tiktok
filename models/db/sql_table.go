package db

import "time"

// 用户数据表

type User struct {
	// gorm.Model
	Id             int    `gorm:"column:id; type:int unsigned; primaryKey;autoIncrement"`
	Username       string `gorm:"column:username; type:varchar(32); not null; unique "`
	Password       string `gorm:"column:password; type:varchar(32); not null;"`
	Nicename       string `gorm:"column:nicename; type:varchar(32); not null;default:'' "`
	Avatar_url     string `gorm:"column:avatar_url; type:varchar(128); not null;default:'' "`
	Follow_count   int    `gorm:"column:follow_count; type:int unsigned; not null;default:0 "`
	Follower_count int    `gorm:"column:follower_count; type:int unsigned; not null;default:0 "`
	Token          string `gorm:"column:token; type:varchar(128); not null;default:'' "`
}

// 用TableName()指定数据库中的表名
func (User) TableName() string {
	return "user"
}

// 视频数据表
type Video struct {
	// gorm.Model
	Id            int       `gorm:"column:id; type:int unsigned; not null;primaryKey;autoIncrement"`
	User_id       int       `gorm:"column:user_id; type:int unsigned; not null"`
	Play_url      string    `gorm:"column:play_url; type:varchar(128); not null"`
	Cover_url     string    `gorm:"column:cover_url; type:varchar(128); not null;default:'' "`
	Like_count    int       `gorm:"column:like_count; type:int unsigned; not null;default:0"`
	Comment_count int       `gorm:"column:comment_count; type:int unsigned; not null;default:0"`
	Title         string    `gorm:"column:title; type:varchar(64); not null;default:'' "`
	Create_time   time.Time `gorm:"column:create_time; type:timestamp; not null;default:current_timestamp "`

	// 设置外键
	User User `gorm:"foreignkey:Id;references:User_id;"`
}

func (Video) TableName() string {
	return "video"
}

// 关注数据表
type Follow struct {
	// gorm.Model
	Id        int `gorm:"column:id; type:int unsigned; not null;primaryKey;autoIncrement"`
	User_id   int `gorm:"column:user_id; type:int unsigned; not null"`
	Follow_id int `gorm:"column:follow_id; type:int unsigned; not null"`

	// 设置外键和级联操作
	User User `gorm:"foreignkey:Id;references:User_id,Follow_id;"`
}

func (Follow) TableName() string {
	return "follow"
}

// 点赞数据表
type Star struct {
	// gorm.Model
	Id       int `gorm:"column:id; type:int unsigned; not null;primaryKey;autoIncrement"`
	User_id  int `gorm:"column:user_id; type:int unsigned; not null"`
	Video_id int `gorm:"column:video_id; type:int unsigned; not null"`

	// 设置外键和级联操作
	User  User  `gorm:"foreignkey:Id;references:User_id;"`
	Video Video `gorm:"foreignkey:Id;references:Video_id;"`
}

func (Star) TableName() string {
	return "star"
}

// 评论数据表
type Comment struct {
	// gorm.Model
	Id           int       `gorm:"column:id; type:int unsigned; not null;primaryKey;autoIncrement"`
	User_id      int       `gorm:"column:user_id; type:int unsigned; not null"`
	Video_id     int       `gorm:"column:video_id; type:int unsigned; not null"`
	Content      string    `gorm:"column:content; type:int unsigned; not null"`
	Content_time time.Time `gorm:"column:content_time; type:timestamp; not null;default:current_timestamp"`

	// 设置外键和级联操作
	User  User  `gorm:"foreignkey:Id;references:User_id;"`
	Video Video `gorm:"foreignkey:Id;references:Video_id;"`
}

func (Comment) TableName() string {
	return "comment"
}

// 聊天数据表
type Chat struct {
	// gorm.Model
	Id       int       `gorm:"column:id; type:int unsigned; not null;primaryKey;autoIncrement"`
	Send_id  int       `gorm:"column:send_id; type:int unsigned; not null"`
	Recv_id  int       `gorm:"column:recv_id; type:int unsigned; not null"`
	Msg      string    `gorm:"column:msg; type:varchar(64); not null"`
	Msg_time time.Time `gorm:"column:msg_time; type:timestamp; not null;default:current_timestamp"`

	// 设置外键和级联操作
	User User `gorm:"foreignkey:Id;references:Send_id,Recv_id;"`
}

func (Chat) TableName() string {
	return "chat"
}
