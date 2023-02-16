package db

type User struct {
	// gorm.Model
	Id       int64  `gorm:"column:user_id; primary_key;"`
	Name     string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Follow   int64  `gorm:"column:follow_count"`
	Follower int64  `gorm:"column:follower_count"`
	Avatar   string `gorm:"column:avatar"`
	TotalFav int64  `gorm:"column:total_favorited"`
	FavCount int64  `gorm:"column:favorite_count"`
}

func (User) TableName() string {
	return "users"
}

type Video struct {
	// gorm.Model
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Title         string `gorm:"column:title;"`
	Author        User   `gorm:"foreignkey:AuthorId"`
}

func (Video) TableName() string {
	return "videos"
}

type Comment struct {
	// gorm.Model
	CommentId int64  `gorm:"column:comment_id; primary_key;"`
	UserId    int64  `gorm:"column:user_id"`
	VideoId   int64  `gorm:"column:video_id"`
	Comment   string `gorm:"column:comment"`
	Time      string `gorm:"column:time"`
}

func (Comment) TableName() string {
	return "comments"
}

type Favorite struct {
	// gorm.Model
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type Relation struct {
	// gorm.Model
	Id       int64 `gorm:"column:relation_id; primary_key;"`
	Follow   int64 `gorm:"column:follow_id"`
	Follower int64 `gorm:"column:follower_id"`
}

func (Relation) TableName() string {
	return "relations"
}

type Message struct {
	// gorm.Model
	MessageId int64  `gorm:"column:message_id; primary_key;"`
	UserId    int64  `gorm:"column:user_id"`
	ToUserId  int64  `gorm:"column:to_user_id"`
	Content   string `gorm:"column:content"`
	Time      int64  `gorm:"column:time"`
}

func (Message) TableName() string {
	return "messages"
}
