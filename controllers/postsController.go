package controllers

import (
	"fmt"
	"go-crud-example/initializers"
	"go-crud-example/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	//get data req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
	}

	//return success

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPost(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	//get post id url
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	//get id
	id := c.Param("id")

	// get data req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find the post
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// get post id url
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Post deleted successfully id: %s", id),
	})

}
