package v1

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"example.com/m/server/models"
	"example.com/m/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerV1) CreatePost(ctx *gin.Context) {
	var req models.CreatePost
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	post, err := h.strg.Post().Create(ctx, &repo.Post{
		ID:     id.String(),
		Title:  req.Title,
		Body:   req.Body,
		UserId: req.UserId,
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.Post{
		ID:       post.ID,
		Title:    post.Title,
		Body:     post.Body,
		UserId:   post.UserId,
		CreateAt: post.CreateAt.Format(time.RFC3339),
	})
}

func (h *handlerV1) UpdatePost(ctx *gin.Context) {
	var req models.UpdatePost
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	id := ctx.Param("id")
	err := h.strg.Post().Update(ctx, &repo.UpdatePost{
		ID:        id,
		Title:     req.Title,
		Body:      req.Body,
		Published: req.Published,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Post not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post updated",
	})
}

func (h *handlerV1) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := h.strg.Post().Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Post not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Post{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		Published: post.Published,
		UserId:    post.UserId,
		CreateAt:  post.CreateAt.Format(time.RFC3339),
	})
}

func (h *handlerV1) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.strg.Post().Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Post not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
