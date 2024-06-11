package controllers

import (
	"context"
	"fmt"
	"io"
	"main_service/db_utils"
	"main_service/models"
	"main_service/post_service"
	"main_service/stats_service"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/types/known/emptypb"
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

// GET /wall/{userId}?offset=100&limit=50
// gets posts with pagination
func GetWall(c *gin.Context) {
	user_id := c.GetString("id")
	var target_id models.UserId
	if err := c.ShouldBindUri(&target_id); err != nil {
		c.JSON(400, gin.H{"error": "wrong user id"})
		return
	}
	var pagination models.Pagination
	c.ShouldBind(&pagination)

	var post_service_pagination post_service.Pagination
	post_service_pagination.Limit = pagination.Limit
	post_service_pagination.Offset = pagination.Offset

	ctx := context.TODO()
	var paginated_request post_service.PaginatedPostRequest
	paginated_request.Pagination = &post_service_pagination
	paginated_request.RequesterId = user_id
	paginated_request.TargetId = target_id.PostId
	stream, err := post_service.Client.GetPaginatedPosts(ctx, &paginated_request)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	var posts []models.Post
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(500, gin.H{"error": "internal server error"})
		}
		posts = append(posts, models.Post{PostTitle: in.PostTitle, PostText: in.PostText})
	}
	c.JSON(200, posts)
}

func sendStats(user_id string, post_id string, author_id string, topic string) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, 0)
	if err != nil {
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(fmt.Sprintf("{\"user_id\":\"%s\",\"post_id\":\"%s\",\"author_id\":\"%s\"}",
			user_id, post_id, author_id))},
	)
	if err != nil {
		return err
	}

	err = conn.Close()
	return err
}

// POST /post/view/{postId}
// add a view to a post
func ViewPost(c *gin.Context) {
	InteractWithPost(c, "views")
}

func LikePost(c *gin.Context) {
	InteractWithPost(c, "likes")
}

// POST /post/like/{postId}
// add a like to a post
func InteractWithPost(c *gin.Context, interactType string) {
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
	var author_id string
	switch post_response.Error {
	case post_service.ErrorEnum_OK:
		c.JSON(200, post_response.Post)
		author_id = post_response.Post.UserId
	case post_service.ErrorEnum_NO_SUCH_POST:
		c.JSON(404, gin.H{"message": "post not found"})
		return
	case post_service.ErrorEnum_ACCESS_DENIED:
		c.JSON(403, gin.H{"message": "user has not access to the post"})
		return
	}

	err = sendStats(user_id, post_id.PostId, author_id, interactType)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func GetPostStats(c *gin.Context) {
	var request stats_service.Id

	var post_id models.PostId
	if err := c.ShouldBindUri(&post_id); err != nil {
		c.JSON(400, gin.H{"error": "wrong post id"})
		return
	}

	request.Id = post_id.PostId
	response, err := stats_service.Client.GetStats(context.TODO(), &request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": post_id.PostId, "likes": response.LikesCount, "views": response.ViewsCount})
}

func GetTopUsers(c *gin.Context) {
	var request emptypb.Empty
	grpc_resp, err := stats_service.Client.GetTopUsers(context.TODO(), &request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Login string
		Count uint64
	}

	var response []Response
	for _, resp := range grpc_resp.Users {
		login, err := db_utils.GetLogin(resp.UserIds)
		if err != nil {
			c.JSON(500, gin.H{"error": "internal server error"})
			return
		}
		response = append(response, Response{Login: login, Count: resp.LikesCount})
	}
	c.JSON(200, response)
}

func GetTopPosts(c *gin.Context) {
	var tp models.Type
	c.ShouldBind(&tp)

	var request stats_service.Type
	request.IsViews = (tp.Type == "views")
	grpc_resp, err := stats_service.Client.GetTopPosts(context.TODO(), &request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		PostId string
		Login  string
		Count  uint64
	}

	var response []Response
	for _, resp := range grpc_resp.Posts {
		login, err := db_utils.GetLogin(resp.AuthorId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		response = append(response, Response{PostId: resp.PostId, Login: login, Count: resp.Count})
	}
	c.JSON(200, response)
}
