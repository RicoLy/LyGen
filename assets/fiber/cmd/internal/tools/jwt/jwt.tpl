package jwt

import (
	"{{.ProjectName}}/cmd/internal/config"
	"{{.ProjectName}}/cmd/internal/tools/cypher"
	"{{.ProjectName}}/cmd/model"
	"time"
)

var Jwt = NewJWT()

type CustomClaims struct {
	PhoneNum   string `json:"phoneNum"`   // 用户id
	Username   string `json:"username"`   // 用户名
	BufferTime int64  `json:"bufferTime"` // 缓冲时间
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(config.GlobalConfig.JWT.SigningKey),
	}
}

// 根据用户信息创建一个token
func (j *JWT) CreateTokenByUserInfo(user *model.User) (string, error) {
	//expireTime := time.Now().Add(time.Duration(2592000) * time.Second)
	//expireTime := time.Now().Add(time.Duration(config.GlobalConfig.JWT.ExpiresTime) * time.Second)
	phoneNum, _ := cypher.EncryptByAes([]byte(user.PhoneNum))
	claims := &CustomClaims{
		PhoneNum:   phoneNum,
		Username:   user.Username,
		BufferTime: config.GlobalConfig.JWT.BufferTime,
		//BufferTime: 2592000,
		StandardClaims: jwt.StandardClaims{
			Issuer: "FiberDemo",
			//ExpiresAt: expireTime.Unix(),
		},
	}
	return j.CreateToken(claims)
}

// 生成token
func (j *JWT) CreateToken(claims *CustomClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(time.Duration(config.GlobalConfig.JWT.ExpiresTime) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	return token.SignedString(j.SigningKey)
}

// 根据tokenStr 解析出Claims
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			phoneNum, _ := cypher.DecryptByAes(claims.PhoneNum)
			claims.PhoneNum = string(phoneNum)
			return claims, nil
		}
	}

	return nil, err
}
