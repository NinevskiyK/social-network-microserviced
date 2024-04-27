package controllers

import (
	"context"
	"io"
	"main_service/models"
	"main_service/post_service"

	"github.com/gin-gonic/gin"
)

// POST /post/create
// creates post, returns id of created post
func CreatePost(c *gin.Context) {
	user_id := c.GetString("id")
	var post_content models.Post
	if err := c.BindJSON(&post_content); err != nil {
		c.JSON(400, gin.H{"error": "wrong post info"})
		return
	}
	ctx := context.TODO()
	post := post_service.Post{}
	post.PostText = post_content.PostText
	post.PostTitle = post_content.PostTitle
	post.UserId = user_id
	post_id, err := post_service.Client.CreatePost(ctx, &post)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": post_id.Id})
}

// PUT /post/update/{postId}
// updates post
func UpdatePost(c *gin.Context) {
	user_id := c.GetString("id")
	var post_id models.PostId
	if err := c.ShouldBindUri(&post_id); err != nil {
		c.JSON(400, gin.H{"error": "wrong post id"})
		return
	}
	var post_content models.Post
	if err := c.BindJSON(&post_content); err != nil {
		c.JSON(400, gin.H{"error": "wrong post info"})
		return
	}
	ctx := context.TODO()
	post := post_service.Post{}
	post.PostText = post_content.PostText
	post.PostTitle = post_content.PostTitle
	post.UserId = user_id
	post.PostId = post_id.PostId
	post_service_error, err := post_service.Client.UpdatePost(ctx, &post)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	switch post_service_error.Error {
	case post_service.ErrorEnum_OK:
		c.JSON(200, gin.H{"message": "updated"})
		return
	case post_service.ErrorEnum_NO_SUCH_POST:
		c.JSON(404, gin.H{"message": "post not found"})
		return
	case post_service.ErrorEnum_ACCESS_DENIED:
		c.JSON(403, gin.H{"message": "user has not access to the post"})
		return
	}
}

// DELETE /post/delete/{postId}
// deletes post
func DeletePost(c *gin.Context) {
	user_id := c.GetString("id")
	var post_id models.PostId
	if err := c.ShouldBindUri(&post_id); err != nil {
		c.JSON(400, gin.H{"error": "wrong post id"})
		return
	}
	ctx := context.TODO()
	var post_request post_service.PostRequest
	post_request.PostId = post_id.PostId
	post_request.RequesterId = user_id
	post_service_error, err := post_service.Client.DeletePost(ctx, &post_request)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	switch post_service_error.Error {
	case post_service.ErrorEnum_OK:
		c.JSON(200, gin.H{"message": "deleted"})
		return
	case post_service.ErrorEnum_NO_SUCH_POST:
		c.JSON(404, gin.H{"message": "post not found"})
		return
	case post_service.ErrorEnum_ACCESS_DENIED:
		c.JSON(403, gin.H{"message": "user has not access to the post"})
		return
	}
}

// GET /post/get/{postId}
// gets post
func GetPost(c *gin.Context) {
	user_id := c.GetString("id")
	var post_id models.PostId
	if err := c.ShouldBindUri(&post_id); err != nil {
		c.JSON(400, gin.H{"error": "wrong post id"})
		return
	}
	ctx := context.TODO()
	var post_request post_service.PostRequest
	post_request.PostId = post_id.PostId
	post_request.RequesterId = user_id
	post_response, err := post_service.Client.GetPost(ctx, &post_request)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	switch post_response.Error {
	case post_service.ErrorEnum_OK:
		c.JSON(200, post_response.Post)
		return
	case post_service.ErrorEnum_NO_SUCH_POST:
		c.JSON(404, gin.H{"message": "post not found"})
		return
	case post_service.ErrorEnum_ACCESS_DENIED:
		c.JSON(403, gin.H{"message": "user has not access to the post"})
		return
	}
}

// GET /wall?offset=100&limit=50
// gets posts with pagination
func GetWall(c *gin.Context) {
	user_id := c.GetString("id")
	var pagination models.Pagination
	c.ShouldBind(&pagination)

	var post_service_pagination post_service.Pagination
	post_service_pagination.Limit = pagination.Limit
	post_service_pagination.Offset = pagination.Offset

	ctx := context.TODO()
	var paginated_request post_service.PaginatedPostRequest
	paginated_request.Pagination = &post_service_pagination
	paginated_request.RequesterId = user_id
	stream, err := post_service.Client.GetPaginatedPosts(ctx, &paginated_request)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	var posts []post_service.Post
	for {
		post, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(500, gin.H{"error": "internal server error"})
		}
		posts = append(posts, *post)
	}
	c.JSON(200, posts)
}
