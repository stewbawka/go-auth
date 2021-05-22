package jwt

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

func CreateToken(ttl time.Duration, data interface{}) (string, error) {
    key, err := jwt.ParseRSAPrivateKeyFromPEM(JWTKeypair.privateKey)
    if err != nil {
        return "", err
    }

    now := time.Now().UTC()

    claims := make(jwt.MapClaims)
    claims["data"] = data
    claims["exp"] = now.Add(ttl).Unix()
    claims["iat"] = now.Unix()
    claims["nbf"] = now.Unix()
    claims["iss"] = "testing@secure.istio.io"


    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    token.Header["kid"] = JwkKid
    tokenStr, err := token.SignedString(key)
    if err != nil {
        return "", err
    }

    return tokenStr, nil
}
