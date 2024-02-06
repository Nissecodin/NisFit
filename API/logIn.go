package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type workout struct {
	Title    string  `json:"title"`
	Creator   string  `json:"author"`
	Description   string   `json:"description"`
    Duration      string   `json:"duration"` 
    Difficulty    string   `json:"difficulty"`
    Equipment     []string `json:"equipment"`
}

var workout = []workout{
	{Title:"1", Creator: "In Search of Lost Time", Description: "Marcel Poust", Duration: 2, Difficulty:"", Equipment},
	{Title:"2", Creator: "The Great Gatsby", Description: "F. Scott Fitzgerald", Duration: 5, Difficulty:"", Equipment,
	{Title:"3", Creator: "War and Peace", Description: "Leo Tolstay", Duration: "", Difficulty:"", Equipment},
}
