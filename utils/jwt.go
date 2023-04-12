package utils

import (
	"crypto/md5"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// three month
const (
	EXPIRE       time.Duration = 72 * time.Hour
	WechatExpire               = 24 * time.Hour
	issuer       string        = "Platform"
)

type Claims struct {
	ID   int64
	Name string
	jwtgo.StandardClaims
}

func CreatJwtKey(id int64) string {
	key := "Jwt:" + strconv.FormatInt(id, 10)
	return key
}

func Messagedigest5(s, salt string) string {
	if (s + salt) == "" {
		s = "123456"
	}
	data := md5.Sum([]byte(s + salt))
	return fmt.Sprintf("%x", data)
}
func DoubleMessagedigest5(s, salt string) string {
	if (s + salt) == "" {
		s = "123456"
	}
	return Messagedigest5(Messagedigest5(s, salt), Messagedigest5(salt, s))
}
func GeneraterJwt(id int64, name, salt string) (string, error) {
	now := time.Now().In(time.Local)
	expire := now.Add(EXPIRE).Unix()
	claim := Claims{
		id,
		name,
		jwtgo.StandardClaims{
			ExpiresAt: expire,
			Issuer:    issuer + name,
		},
	}
	tokenclaim := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claim)
	jwt, err := tokenclaim.SignedString([]byte(salt))
	return jwt, err
}
