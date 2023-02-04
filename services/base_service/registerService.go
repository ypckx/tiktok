package base_service

import (
	"tinyTiktok/models"
)

/*
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string"
}
*/

// type PostRegisterUser struct {
// 	username string
// 	password string
// 	data     *models.RegisterResp
// }

func UserReigsterService(username string, password string) (retResp models.RegisterResp, retErr error) {

	idAndToken, err := models.UserReigsterDao(models.DB.Statement.Context, username, password)
	if err != nil {
		retResp = models.RegisterResp{
			Response: models.Response{StatusCode: 1, StatusMsg: "User already exist"},
		}
		retErr = err
		return retResp, retErr
	}

	retResp = models.RegisterResp{
		Response: models.Response{StatusCode: 0},
		UserId:   idAndToken.UserId,
		Token:    idAndToken.Token,
	}

	return retResp, nil
}
