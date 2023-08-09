package token_verify_middleware

import (
	"mentor_app/user_system/dto"
	"mentor_app/user_system/errcode"
	"mentor_app/user_system/initialize"
	"mentor_app/user_system/middleware/log"
	"mentor_app/user_system/vo"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var jwtSecret = []byte(initialize.GLOBAL_CONFIG.Cert.JwtSecret)

func VerifyToken(c *gin.Context) {
	jwtToken := c.GetHeader("authToken")

	err, claim := getClaimFromJwt(jwtToken)
	if err != nil {
		c.JSON(200, vo.NewErrorResponse(errcode.TokenDecodeError))
		c.Abort()
		return
	}

	isExpired := !checkExpireTime(claim.ExpireTime)
	if isExpired {
		c.JSON(200, vo.NewErrorResponse(errcode.TokenExpiredError))
		c.Abort()
		return
	}

	c.Set("jwtClaim", claim)
	c.Next()
}

func VerifyTokenWithApi(userToken string) (error, *dto.JwtClaim) {
	err, claim := getClaimFromJwt(userToken)
	if err != nil {
		return errors.New("verify fail"), nil
	}

	isExpired := !checkExpireTime(claim.ExpireTime)
	if isExpired {
		return errors.New("token expired"), nil
	}

	return nil, claim
}

func getClaimFromJwt(token string) (error, *dto.JwtClaim) {
	var claim *dto.JwtClaim
	tokenClaims, err := jwt.ParseWithClaims(token, &dto.JwtClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		log.Logger.Error(err)
		return err, nil
	}

	if claims, ok := tokenClaims.Claims.(*dto.JwtClaim); ok {
		claim = claims
	}

	return nil, claim
}

func checkExpireTime(tokenTime string) bool {
	format := "2006/1/2 15:04:05"
	exprieTime, _ := time.Parse(format, tokenTime)

	return exprieTime.After(time.Now().UTC())
}
