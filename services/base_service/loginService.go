package base_service

import (
	"tinyTiktok/models"
)

// 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
func UserLoginService(username string, password string) (retResp models.LoginResp, retErr error) {
	// userToken := models.GenerateToken(username, password)
	// _,ex models.TokenToIdMap[userToken]
	idAndToken, err := models.UserLoginDao(models.DB.Statement.Context, username, password)
	if err != nil {
		retResp = models.LoginResp{
			Response: models.Response{StatusCode: 1, StatusMsg: err.Error()},
		}
		retErr = err
		return retResp, retErr
	}

	retResp = models.LoginResp{
		Response: models.Response{StatusCode: 0},
		UserId:   idAndToken.UserId,
		Token:    idAndToken.Token,
	}
	return retResp, nil
}
