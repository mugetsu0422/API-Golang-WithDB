package controller

import (
	"API-Golang-WithDB/model"
	"API-Golang-WithDB/storage"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

// GET API
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
	student := &model.Students{}

	// Check if exist
	id, _ := strconv.Atoi(c.Param("id"))
	if err := db.Take(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Not found student")
	}

	return c.JSON(http.StatusOK, student)
}

// POST API
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

// PUT API
func UpdateStudent(c echo.Context) error {
	db := storage.GetDBInstance()
	student := &model.Students{}

	// Check if exist
	id, _ := strconv.Atoi(c.Param("id"))
	if err := db.Take(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Not found student")
	}

	if err := c.Bind(student); err != nil {
		return err
	}
	student.Id = id

	db.Save(student)
	return c.JSON(http.StatusOK, student)
}

// DELETE API
func DeleteStudent(c echo.Context) error {
	db := storage.GetDBInstance()
	student := &model.Students{}

	// Check if exist
	id, _ := strconv.Atoi(c.Param("id"))
	if err := db.Take(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Not found student")
	}

	db.Delete(student, id)
	return c.JSON(http.StatusOK, student) 
}


// Forget Password API 
// Method: POST
// Path: /login
// Content-Type: JSON
// Form:
// {
//      "username": "username",
//      "email": "email"
// }
func ForgetPassword(c echo.Context) error {
	db := storage.GetDBInstance()
	user := &model.User{}

	if err := c.Bind(user); err != nil {
		return err;
	}

	if result := db.Where("username = ? AND email = ?", user.Username, user.Email).Take(&user); 
	result.Error != nil {
		return result.Error;
	}

	return c.JSON(http.StatusOK, user)
}

// Like Password API 
// Method: POST
// Path: /like
// Content-Type: JSON
// Form:
// {
//      "user_id": "user_id",
//      "song_id": "song_id"
// }
func LikeSong(c echo.Context) error {
	db := storage.GetDBInstance()
	like_song := &model.Liked_song{}

	if err := c.Bind(like_song); err != nil {
		return err;
	}

	if result := db.Create(&like_song); 
	result.Error != nil {
		return result.Error;
	}

	return c.JSON(http.StatusCreated, like_song)
}

func DislikeSong(c echo.Context) error {
	db := storage.GetDBInstance()
	like_song := &model.Liked_song{}

	if err := c.Bind(like_song); err != nil {
		return err;
	}

	if result := db.Where("user_id = ? AND song_id = ?", like_song.User_id, like_song.Song_id).Delete(&like_song); 
	result.Error != nil {
		return result.Error;
	}
	return c.JSON(http.StatusOK, like_song) 
}

// Get Song info API
// Method: GET
// Path: /song
// Content-Type: JSON
// Form:
// {
//      "song_id": "song_id"
// }
func GetSong(c echo.Context) error {
	db := storage.GetDBInstance()
	song := &model.Song{}

	if err := c.Bind(song); err != nil {
		return err;
	}
	result := db.Take(&song, "song_id = ?", song.Song_id);
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Song not found");
	}
	return c.JSON(http.StatusOK, song) 
}