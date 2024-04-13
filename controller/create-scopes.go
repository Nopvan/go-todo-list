package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

type CreateScopesRequest struct {
	Name string `json:"name"`
}

func CreateScopes(e *echo.Echo, db *sql.DB) {

	e.POST("/scopes", func(ctx echo.Context) error {

		var request CreateScopesRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		//untuk mengambil row di table scopes
		row := db.QueryRow("SELECT id FROM scopes WHERE name = ?", request.Name)
		if row.Err() != nil {
			//kalau gak nemu baris akan return error
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		//kondisi untuk ngecheck datanya sudah ada atau belum
		var retrivedId int
		err := row.Scan(&retrivedId)
		//jika ada errornya berarti datanya sudah ada/double
		if err == nil {
			return ctx.String(http.StatusBadRequest, "Duplicate scope found")
		}
		//jika gak ada error dan diluar dari error scope return internal server error
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		//ini versi codingan rapihnya
		/*
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return ctx.String(http.StatusBadRequest, "Duplicate scope found")
				}
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
		*/

		_, err = db.Exec(
			"INSERT INTO scopes (name) VALUES (?)",
			request.Name,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")

	})
}
