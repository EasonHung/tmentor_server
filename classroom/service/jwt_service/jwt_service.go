package jwt_service

import (
	"mentor/classroom/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func GenerateClasstoken(mentorId string, classroomId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["mentorId"] = mentorId
	atClaims["classroomId"] = classroomId
	atClaims["expireTime"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.GLOBAL_CONFIG.Secret.JwtKey))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return token, nil
}

func VerifyTokenAndReturnClassroomId(classroomToken string) (error, string) {
	token, err := jwt.Parse(classroomToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GLOBAL_CONFIG.Secret.JwtKey), nil
	})
	if err != nil {
		return err, ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		classroomId, ok := claims["classroomId"].(string)
		if !ok {
			return err, ""
		}
		expireTime, ok := claims["expireTime"].(float64)
		if !ok {
			return err, ""
		}
		if int64(expireTime) < time.Now().Unix() {
			return errors.New("token expire time is invalid"), ""
		}

		return nil, classroomId
	}
	return errors.WithStack(err), ""
}