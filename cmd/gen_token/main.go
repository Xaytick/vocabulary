package main

import (
	"fmt"
	"time"
	"vocabulary/internal/consts"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Id       uint   `json:"Id"`
	Username string `json:"Username"`
	jwt.RegisteredClaims
}

func main() {
	// 设置你想要的用户信息
	claims := &jwtClaims{
		Id:       1, // 用户ID
		Username: "Wang Xiang", // 用户名
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(consts.JwtKey))
	if err != nil {
		fmt.Println("生成token失败:", err)
		return
	}

	fmt.Println("生成的token:")
	fmt.Println(tokenString)
	
	fmt.Println("\n使用curl测试:")
	fmt.Printf("curl -H \"Authorization: Bearer %s\" http://127.0.0.1:8000/v1/account/info\n", tokenString)
}