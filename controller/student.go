package controller

import (
	"API-Golang-WithDB/model"
	"API-Golang-WithDB/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents() 
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	db := storage.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}