package token_verification_service

// todo: seperate this service into another independent micro service

import (
	"mentor_app/user_system/dto"
	"mentor_app/user_system/initialize"
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(initialize.GLOBAL_CONFIG.Cert.JwtSecret)

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
