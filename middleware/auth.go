package middleware

import (
	"code/api"
	"code/global"
	"code/global/constants"
	"code/model"
	"code/service"
	"code/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ErrCodeInvalidToken    = 10401
	ErrCodeTokenParse      = 10402
	ErrCodeTokenNotMatched = 10403
	ErrCodeTokenExpired    = 10404
	ErrCodeTokenRenew      = 10405

	TokenName          = "Authorization"
	TokenPrefix        = "Bearer "
	RenewTokenDuration = 10 * 60 * time.Second
)

func tokenErr(c *gin.Context, code int) {
	api.Fail(c, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "Invalid Token",
	})
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(TokenName)
		if token == "" || !strings.HasPrefix(token, TokenPrefix) {
			tokenErr(c, ErrCodeInvalidToken)
			return
		}

		// 判断token格式
		token = strings.TrimPrefix(token, TokenPrefix)
		iJwtCustClaims, err := utils.ParseToken(token)
		nUserID := iJwtCustClaims.ID
		if err != nil || nUserID == 0 {
			tokenErr(c, ErrCodeTokenParse)
			return
		}

		// 判断token是否相同
		stUserID := strconv.Itoa(int(nUserID))
		stRedisUserIDKey := strings.Replace(constants.LoginUserTokenRedisKey, "{ID}", stUserID, -1)
		stRedisToken, err := global.RedisClient.Get(stRedisUserIDKey)
		if err != nil || stRedisToken != token {
			tokenErr(c, ErrCodeTokenNotMatched)
			return
		}

		// 判断token是否过期
		nTokenExpireDuration, err := global.RedisClient.GetExpireDuration(stRedisUserIDKey)
		if err != nil || nTokenExpireDuration < 0 {
			tokenErr(c, ErrCodeTokenExpired)
			return
		}

		// token续期
		if nTokenExpireDuration.Seconds() < RenewTokenDuration.Seconds() {
			stNewToken, err := service.GenerateAndCacheLoginUserTokenToRedis(nUserID, iJwtCustClaims.Name)
			if err != nil {
				tokenErr(c, ErrCodeTokenRenew)
				return
			}
			c.Header("token", stNewToken)
		}

		//iUser, err := dao.NewUserDao().GetUserByID(nUserID)
		//if err != nil {
		//	return
		//}
		//c.Set(constants.LoginUser, iUser)
		c.Set(constants.LoginUser, model.LoginUser{
			ID:   nUserID,
			Name: iJwtCustClaims.Name,
		})
		c.Next()
	}
}
