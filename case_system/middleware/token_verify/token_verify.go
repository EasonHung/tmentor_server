package token_verify_middleware

import (
	"case_system/config"
	"case_system/dto"
	"case_system/middleware/log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(config.GLOBAL_CONFIG.Cert.JwtSecret)

func VerifyToken(c *gin.Context) {
	jwtToken := c.GetHeader("authToken")

	err, claim := getClaimFromJwt(jwtToken)
	if err != nil {
		c.JSON(500, err)
		c.Abort()
		return
	}

	isExpired := !checkExpireTime(claim.ExpireTime)
	if isExpired {
		c.JSON(403, "token expired")
		c.Abort()
		return
	}

	c.Set("jwtClaim", claim)
	c.Next()
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

	return exprieTime.After(time.Now())
}
