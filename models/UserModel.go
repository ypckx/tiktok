package models

import (
	"context"
	"fmt"
	"tinyTiktok/models/db"
)

func UserReigsterDao(ctx context.Context, username string, password string) (IdAndToken, error) {
	var user db.User
	conn := DB.WithContext(ctx).Where("username = ?", username).Find(&user)
	if conn.RowsAffected == 0 {
		return IdAndToken{}, conn.Error
	}

	geneToken := GenerateToken(username, password)
	user.Password = password
	user.Username = username
	user.Token = geneToken

	if err := DB.Create(&user).Error; err != nil {
		fmt.Println("UserReigsterDao error!")
		return IdAndToken{}, err
	}

	TokenToIdMap[geneToken] = user.Id
	return IdAndToken{UserId: int64(user.Id), Token: geneToken}, nil
}

// 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
func UserLoginDao(ctx context.Context, username string, password string) (IdAndToken, error) {
	var user db.User
	if err := DB.WithContext(ctx).Where("username = ? AND password = ?", username, password).Find(&user).Error; err != nil {
		return IdAndToken{}, err
	}
	return IdAndToken{UserId: int64(user.Id), Token: user.Token}, nil
}

// ============== dao
