package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func AssignScope(e *echo.Echo, db *sql.DB) {

	e.POST("/users/:userId/scopes/:scopeId/assign", func(ctx echo.Context) error {
		// user := ctx.Get("USER").(models.AuthClaimJWT)
		// fmt.Println(user.UserEmail)

		userId := ctx.Param("userId")
		scopeId := ctx.Param("scopeId")

		//untuk mengambil row di table user_scopes
		row := db.QueryRow("SELECT id FROM user_scopes WHERE user_id = ? AND scope_id = ?", userId, scopeId)
		if row.Err() != nil {
			//kalau gak nemu baris akan return error
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		//kondisi untuk ngecheck datanya sudah ada atau belum
		var retrivedId int
		err := row.Scan(&retrivedId)
		//jika ada errornya berarti datanya sudah ada/double
		if err == nil {
			return ctx.String(http.StatusBadRequest, "Duplicate assignment found")
		}
		//jika gak ada error dan diluar dari error user scopes return internal server error
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		//ini versi codingan rapihnya
		/*
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return ctx.String(http.StatusBadRequest, "Duplicate assignment found")
				}
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
		*/

		_, err = db.Exec(
			"INSERT INTO user_scopes (user_id, scope_id) VALUES (?, ?)",
			userId,
			scopeId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")

	})
}
