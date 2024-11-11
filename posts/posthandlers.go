package posts

import (
	"ChatTryCreate/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var SamplePost = []models.Post{
	{Id: 0, Title: "Title_0", Body: "Body post _ 0", Author: "Me", Time: "11.11.2024"},
	{Id: 1, Title: "Title_1", Body: "Body post _ 1", Author: "You", Time: "10.11.2024"},
	{Id: 2, Title: "Title_2", Body: "Body post _ 2", Author: "We", Time: "09.11.2024"},
}

func ShowPosts() []models.Post {
	return SamplePost
}

func ShowPostsOnPage(c *gin.Context) {
	posts := ShowPosts()
	c.HTML(http.StatusOK,
		"posts.html",
		gin.H{"payload": posts})
}

func ShowPostsJson(c *gin.Context) {
	posts := ShowPosts()
	c.JSON(http.StatusOK, posts)
}

func ShowPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Post{}"})
}
