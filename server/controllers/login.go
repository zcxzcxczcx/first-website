package controllers

import (
	"context"
	"fmt"
	"time"

	"first_website/proto"
	"first_website/server/services"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

//自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

// Login implements
func (s *Server) Login(ctx context.Context, in *proto.LoginReq) (
	resp *proto.LoginResp, err error) {
	name := in.Name
	pwd := in.Password
	maxAge := 60 * 60 * 24
	customClaims := &CustomClaims{
		UserId: 1, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
			Issuer:    name,                                                       // 非必须，也可以填充用户名，
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n", tokenString)
	// 把用户数据存入mongo
	srv := new(services.LoginSrv)
	err = srv.Save(services.User{
		Name:     name,
		Password: pwd,
	})
	if err != nil {
		return
	}
	resp = new(proto.LoginResp)
	resp.Token = tokenString
	return
}

// // IsLogin
// func IsLogin(ctx context.Context, tokenString string) (*CustomClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(SECRETKEY), nil
// 	})
// 	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
// 		return claims, nil
// 	} else {
// 		return nil, err
// 	}
// }
