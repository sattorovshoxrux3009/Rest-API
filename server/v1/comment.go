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

func (h *handlerV1) CreateComment(ctx *gin.Context) {
	var req models.CreateComment
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
	comment, err := h.strg.Comment().Create(ctx, &repo.Comment{
		ID:     id.String(),
		Body:   req.Body,
		PostId: req.PostId,
		UserId: req.UserId,
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.Comment{
		ID:       comment.ID,
		Body:     comment.Body,
		UserId:   comment.UserId,
		PostId:   comment.PostId,
		CreateAt: comment.CreateAt.Format(time.RFC3339),
	})
}

func (h *handlerV1) UpdateComment(ctx *gin.Context) {
	var req models.UpdateComment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	id := ctx.Param("id")
	err := h.strg.Comment().Update(ctx, &repo.UpdateComment{
		ID:   id,
		Body: req.Body,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Comment not found",
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
		"message": "Comment updated",
	})
}

func (h *handlerV1) GetComment(ctx *gin.Context) {
	id := ctx.Param("id")
	comment, err := h.strg.Comment().Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Comment not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Comment{
		ID:       comment.ID,
		Body:     comment.Body,
		PostId:   comment.PostId,
		UserId:   comment.UserId,
		CreateAt: comment.CreateAt.Format(time.RFC3339),
	})
}

func (h *handlerV1) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.strg.Comment().Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Comment not found",
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
		"message": "Comment deleted",
	})
}
