package handlers

import (
	"net/http"

	"blog-backend/config"
	"blog-backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var postInput struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&postInput); err != nil {
		RespondError(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	userToken, exists := c.Get("user")
	if !exists {
		RespondError(c, http.StatusUnauthorized, "user not found in token", nil)
	}

	claims := userToken.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	post := models.Post{
		Title:   postInput.Title,
		Content: postInput.Content,
		UserID:  userID,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		RespondError(c, http.StatusInternalServerError, "failed to create post", err)
		return
	}
	RespondSuccess(c, "post created", gin.H{"post_id": post.ID})
}

func CreateComment(c *gin.Context) {
	var commentInput struct {
		PostID  uint   `json:"post_id" binding:"required"`
		Content string `json:"content" binding:"required"` // 拼写修正
	}

	if err := c.ShouldBindJSON(&commentInput); err != nil {
		RespondError(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	userToken, exists := c.Get("user")
	if !exists {
		RespondError(c, http.StatusUnauthorized, "user not found in token", nil)
		return
	}

	claims := userToken.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	comment := models.Comment{
		Content: commentInput.Content, // 修正了原来的 Cotent 拼写
		UserID:  userID,
		PostID:  commentInput.PostID,
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		RespondError(c, http.StatusInternalServerError, "failed to create comment", err)
		return
	}

	RespondSuccess(c, "comment created", gin.H{"comment_id": comment.ID})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := config.DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	postID := c.Param("id")

	var post models.Post

	if err := config.DB.Preload("User").First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func GetComments(c *gin.Context) {
	postID := c.Param("post_id")

	var comments []models.Comment
	if err := config.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch comments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func UpdatePost(c *gin.Context) {
	postID := c.Param("id")

	var postInput struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&postInput); err != nil {
		RespondError(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	userToken, _ := c.Get("user")
	claims := userToken.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		RespondError(c, http.StatusNotFound, "post not found", err)
		return
	}

	if post.UserID != userID {
		RespondError(c, http.StatusForbidden, "you are not the author", nil)
		return
	}

	post.Title = postInput.Title
	post.Content = postInput.Content

	if err := config.DB.Save(&post).Error; err != nil {
		RespondError(c, http.StatusInternalServerError, "failed to update post", err)
		return
	}

	RespondSuccess(c, "post updated", gin.H{"post_id": post.ID})
}

func DeletePost(c *gin.Context) {
	postID := c.Param("id")

	userToken, exists := c.Get("user")
	if !exists {
		RespondError(c, http.StatusUnauthorized, "user not found in token", nil)
		return
	}
	claims := userToken.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		RespondError(c, http.StatusNotFound, "post not found", err)
		return
	}

	if post.UserID != userID {
		RespondError(c, http.StatusForbidden, "you are not the author", nil)
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		RespondError(c, http.StatusInternalServerError, "failed to delete post", err)
		return
	}

	RespondSuccess(c, "post deleted", gin.H{"post_id": post.ID})
}
