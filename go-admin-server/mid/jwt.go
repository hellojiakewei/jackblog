package mid

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jackblog/model/response"
	"strings"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义过期时间
var TokenExpire = time.Hour * 2

// 定义secret
var Myscret = []byte("helloJack")

/**
 * @Description: 生成jwt
 * @param username
 * @return string
 * @return error
 */
func GenToken(username string) (string, error) {
	c := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpire).Unix(), // 过期时间
			Issuer:    "jacketed",                         // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Myscret)
}

// 解析jwt
func ParseToken(UserToken string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(UserToken, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Myscret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效token")
}

// 更新token
//func  RefreshToken(tokenString string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return Myscret, nil
//	})
//	if err != nil {
//		return "", err
//	}
//	if claims, ok := token.Claims.(MyClaims); ok && token.Valid {
//		jwt.TimeFunc = time.Now
//		claims.StandardClaims.ExpiresAt = time.Now().Unix() + 60*60*24*7
//		return jwt.CreateToken(*claims)
//	}
//	return "", TokenInvalid
//}

// token 中间件
func JWTAuthMIddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			response.FailWithMessage("token为必传项", c)
			c.Abort()
			return
		}
		// 按照空格分隔 获取 Bearer 后边的token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("请求头中 auth格式错误", c)
			c.Abort()
			return
		}
		info, err := ParseToken(parts[1])
		if err != nil {
			response.FailWithMessage("无效token", c)
			c.Abort()
			return
		}
		if time.Now().Unix() > info.ExpiresAt {
			response.FailWithMessage("token登录失效,请重新登录", c)
			c.Abort()
			return
		}

		c.Set("userInfo", info)
		c.Next()
	}
}
