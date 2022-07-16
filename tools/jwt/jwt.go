package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2 // 过期时间
var mySecret = []byte("纳兰嫣然")             // 签名

// MyClaims 自定义结构体，并内嵌jwt.StandardClaims
// jwt.StandardClaims只包含官方字段
// 需要自定义需要的字段
type MyClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(userId int64, username string) (string, error) {
	// 当前时间
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(TokenExpireDuration)
	//   签发人
	issuer := "bluebell"
	//	 赋值给结构体
	claims := MyClaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 转成纳秒
			Issuer:    issuer,
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==》 头部，载荷，签证
	toke, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
	return toke, err
}

// ParseToken 验证token
func ParseToken(tokenStr string) (*MyClaims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil { // 解析token失败
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
