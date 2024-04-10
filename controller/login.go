package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"go-todo-list/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func Login(e *echo.Echo, db *sql.DB) {

	e.POST("/auth/login", func(ctx echo.Context) error {

		var request LoginRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		//mengakses db buat cari email yang di input
		row := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", request.Email)
		//check error, karna errornya dari query row di embed di row checknya kayak gini
		if row.Err() != nil {
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		var retrivedId int
		var retrivedName, retrivedEmail, retrivedPassword string

		//untuk mengambil data dari query / nerima variabel2 yang di inputin
		err := row.Scan(&retrivedId, &retrivedName, &retrivedEmail, &retrivedPassword)
		if err != nil {
			//untuk check apakah email ada atau tidak (jika no rows brarti tidak ada pake package errors trus pake method Is)
			if errors.Is(err, sql.ErrNoRows) {
				return ctx.String(http.StatusUnauthorized, "email is not registered")
			}
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		//untuk compare password yang sudah di hash di database dengan password yang di inputkan / request dengan method CompareHashAndPassword
		err = bcrypt.CompareHashAndPassword([]byte(retrivedPassword), []byte(request.Password))
		if err != nil {
			return ctx.String(http.StatusUnauthorized, "wrong password")
		}

		//diambil dari struct authclaimjwt dan isi data datanya pake data yang diambil dari database
		tokenClaim := models.AuthClaimJWT{
			UserId:    retrivedId,
			UserName:  retrivedName,
			UserEmail: retrivedEmail,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
		//token ini berupa objek dimana bisa akses beberapa method, kalau kita ingin mempublish / mengkonversi jadi string pake signedstring
		//dan si signedstring menerima secret key(better di taro di tempat aman contoh : env)
		tokenStr, err := token.SignedString([]byte("TEST"))
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		response := LoginResponse{
			AccessToken: tokenStr,
		}

		return ctx.JSON(http.StatusOK, response)

	})
}
