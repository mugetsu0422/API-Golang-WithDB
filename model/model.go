package model

import (
	
)

type Students struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	User_id      string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
	Gender       int    `json:"gender"`
	Phone_number string `json:"phone_number"`
}

type Song struct {
	Song_id     string `json:"song_id"`
	Uploaded_by string `json:"uploaded_by"`
	Performer   string `json:"performer"`
	Title       string `json:"title"`
	Play_count  int64  `json:"play_count"`
	Liked_count int64  `json:"liked_count"`
	Status      string `json:"status"`
}

type Session_key struct {
	Key string `json:"key"`
	User_id string `json:"user_id"`
	Time int64 `json:"time"`
	Device string `json:"device"`
}

type Role_permission struct {
	Role_id string `json:"role_id"`
	Permission string `json:"permission"`
	Status string `json:"status"`
}

type Playlist_content struct {
	Playlist_id string `json:"playlist_id"`
	Song_id string `json:"song_id"`
}

type Playlist struct {
	Playlist_id string `json:"playlist_id"`
	Created_by string `json:"created_by"`
	Type string `json:"type"`
	Title string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Published_date int64 `json:"published_date"`
	Created_date int64 `json:"created_date"`
}

type Password_salt struct {
	User_id string `json:"user_id"`
	Salt string `json:"salt"`
}

type Liked_song struct {
	User_id string `json:"user_id"`
	Song_id string `json:"song_id"`
}