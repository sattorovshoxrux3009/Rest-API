package v1

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"example.com/m/server/models"
	"example.com/m/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerV1) CreateSavedPost(ctx *gin.Context) {
	var req models.CreateSavedPost
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
	saved_post, err := h.strg.Saved_post().Create(ctx, &repo.SavedPost{
		ID:     id.String(),
		PostID: req.PostID,
		UserID: req.UserID,
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.SavedPost{
		ID:     saved_post.ID,
		PostID: saved_post.PostID,
		UserID: saved_post.UserID,
	})
}

func (h *handlerV1) GetSavedPost(ctx *gin.Context) {
	id := ctx.Param("id")
	saved_post, err := h.strg.Saved_post().Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Saved post not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error we got :(",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.SavedPost{
		ID:     saved_post.ID,
		PostID: saved_post.PostID,
		UserID: saved_post.UserID,
	})
}

func (h *handlerV1) DeleteSavedPost(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.strg.Saved_post().Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Saved post not found",
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
		"message": "Saved post deleted",
	})
}
