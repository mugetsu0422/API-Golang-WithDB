package controller

import (
	"API-Golang-WithDB/model"
	"API-Golang-WithDB/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllStudents(c echo.Context) error {
	db := storage.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Not found student")
	}

	return c.JSON(http.StatusOK, students)
}

func GetStudent(c echo.Context) error {
	db := storage.GetDBInstance()
	student := model.Students{}

	id, _ := strconv.Atoi(c.Param("id"))
	if err := db.Find(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Not found student")
	}

	return c.JSON(http.StatusOK, student)
}

func CreateStudent(c echo.Context) error {
	db := storage.GetDBInstance()
	temp := &model.Students{}
	db.Last(&temp)
	student := &model.Students{
		Id: temp.Id + 1,
	}

	if err := c.Bind(student); err != nil {
		return err
	}

	db.Create(&student)
	return c.JSON(http.StatusCreated, student)
}

func GetRepoStudents() ([]model.Students, error) {
	db := storage.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}