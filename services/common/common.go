package common

// type ComResponse models.Response
// type RespVideo models.Video
// type RespUser models.User
// type RespComment models.Comment
// type RespMessage models.Message
// type RespMessageSendEvent models.MessageSendEvent
// type RespMessagePushEvent models.MessagePushEvent

/*****************************
定义每个路由要响应，返回的结构体
*****************************/

// /*
// 	{
// 	    "status_code": 0,
// 	    "status_msg": "string",
// 	    "user_id": 0,
// 	    "token": "string"
// 	}
// */
// type RegisterResp struct {
// 	models.Response
// 	UserId int64  `json:"user_id,omitempty"`
// 	Token  string `json:"token,omitempty"`
// }

// /*
// 	{
// 	    "status_code": 0,
// 	    "status_msg": "string",
// 	    "user_id": 0,
// 	    "token": "string"
// 	}
// */
// type LoginResp struct {
// 	models.Response
// 	UserId int64  `json:"user_id,omitempty"`
// 	Token  string `json:"token,omitempty"`
// }

// /*
// 	{
// 	    "status_code": 0,
// 	    "status_msg": "string",
// 	    "user": {
// 	        "id": 0,
// 	        "name": "string",
// 	        "follow_count": 0,
// 	        "follower_count": 0,
// 	        "is_follow": true
// 	    }
// 	}
// */
// type UserInfoResp struct {
// 	models.Response
// 	UserInfo models.User `json:"user"`
// }

// /*
// 	{
// 	    "status_code": 0,
// 	    "status_msg": "string"
// 	}
// */
// type PublishActionResp struct {
// 	models.Response
// }

// /*
// {
//     "status_code": 0,
//     "status_msg": "string",
//     "video_list": [
//         {
//             "id": 0,
//             "author": {
//                 "id": 0,
//                 "name": "string",
//                 "follow_count": 0,
//                 "follower_count": 0,
//                 "is_follow": true
//             },
//             "play_url": "string",
//             "cover_url": "string",
//             "favorite_count": 0,
//             "comment_count": 0,
//             "is_favorite": true,
//             "title": "string"
//         }
//     ]
// }
// */

// type PublishListResp struct {
// 	models.Response
// 	VideoList []models.Video `json:"video_list,omitempty"`
// 	NextTime  int64          `json:"next_time,omitempty"`
// }

// /*
// 	{
// 	    "status_code": 0,
// 	    "status_msg": "string"
// 	}
// */
// type FavoriteActionResp struct {
// 	models.Response
// }

// /*
// 	{
// 	    "status_code": "string",
// 	    "status_msg": "string",
// 	    "video_list": [
// 	        {
// 	            "id": 0,
// 	            "author": {
// 	                "id": 0,
// 	                "name": "string",
// 	                "follow_count": 0,
// 	                "follower_count": 0,
// 	                "is_follow": true
// 	            },
// 	            "play_url": "string",
// 	            "cover_url": "string",
// 	            "favorite_count": 0,
// 	            "comment_count": 0,
// 	            "is_favorite": true,
// 	            "title": "string"
// 	        }
// 	    ]
// 	}
// */
// type FavoriteListResp struct {
// 	models.Response
// 	VideoList []models.Video `json:"video_list,omitempty"`
// }

// /*
// {
//     "status_code": 0,
//     "status_msg": "string",
//     "comment": {
//         "id": 0,
//         "user": {
//             "id": 0,
//             "name": "string",
//             "follow_count": 0,
//             "follower_count": 0,
//             "is_follow": true
//         },
//         "content": "string",
//         "create_date": "string"
//     }
// }
// */

// type CommentActionResp struct {
// 	models.Response
// 	Comment models.Comment `json:"comment,omitempty"`
// }

// /*
// {
//     "status_code": 0,
//     "status_msg": "string",
//     "comment_list": [
//         {
//             "id": 0,
//             "user": {
//                 "id": 0,
//                 "name": "string",
//                 "follow_count": 0,
//                 "follower_count": 0,
//                 "is_follow": true
//             },
//             "content": "string",
//             "create_date": "string"
//         }
//     ]
// }
// */

// type CommentListResp struct {
// 	models.Response
// 	CommentList []models.Comment `json:"comment_list,omitempty"`
// }

// /*
// 	{
// 	    "status_code": 0,
// 	    "status_msg": "string"
// 	}
// */
// type RelationActionResp struct {
// 	models.Response
// }

// /*
// 	{
// 	    "status_code": "string",
// 	    "status_msg": "string",
// 	    "user_list": [
// 	        {
// 	            "id": 0,
// 	            "name": "string",
// 	            "follow_count": 0,
// 	            "follower_count": 0,
// 	            "is_follow": true
// 	        }
// 	    ]
// 	}
// */
// type FollowListResp struct {
// 	models.Response
// 	UserList []models.User `json:"user_list,omitempty"`
// }

// /*
// 	{
// 	    "status_code": "string",
// 	    "status_msg": "string",
// 	    "user_list": [
// 	        {
// 	            "id": 0,
// 	            "name": "string",
// 	            "follow_count": 0,
// 	            "follower_count": 0,
// 	            "is_follow": true
// 	        }
// 	    ]
// 	}
// */
// type FollowerListResp struct {
// 	models.Response
// 	UserList []models.User `json:"user_list,omitempty"`
// }

// /*
// 	{
// 	    "status_code": "string",
// 	    "status_msg": "string",
// 	    "user_list": [
// 	        {
// 	            "id": 0,
// 	            "name": "string",
// 	            "follow_count": 0,
// 	            "follower_count": 0,
// 	            "is_follow": true
// 	        }
// 	    ]
// 	}
// */
// type FriendListResp struct {
// 	models.Response
// 	UserList []models.User `json:"user_list,omitempty"`
// }

// /*
// 	{
// 	    "status_code": 0,
// 	    "status_msg": "string"
// 	}
// */
// type MessageActionResp struct {
// 	models.Response
// }

// /*
// 	{
// 	    "status_code": "string",
// 	    "status_msg": "string",
// 	    "message_list": [
// 	        {
// 	            "id": 0,
// 	            "content": "string",
// 	            "create_time": "string"
// 	        }
// 	    ]
// 	}
// */
// type MessageChatResp struct {
// 	models.Response
// 	MessageList []models.Message `json:"message_list,omitempty"`
// }
