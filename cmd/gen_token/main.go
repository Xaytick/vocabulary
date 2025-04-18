package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Id       uint   `json:"Id"`
	Username string `json:"Username"`
	jwt.RegisteredClaims
}

func main() {
	ctx := context.Background()
	claims := &jwtClaims{
		Id:       1, // 用户ID
		Username: "Wang Xiang", // 用户名
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := g.Cfg().MustGet(ctx, "server.jwt.secret").String()
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		fmt.Println("生成token失败:", err)
		return
	}

	fmt.Println("生成的token:")
	fmt.Println(tokenString)
	
}