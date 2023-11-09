package router

import (
	"net/http"

	"github.com/erbalo/hexago/internal/app/card"
	"github.com/erbalo/hexago/internal/app/domain"
	"github.com/gin-gonic/gin"
)

type CardRouter struct {
	cardService card.Service
}

func NewCardEntry(cardService card.Service) *CardRouter {
	return &CardRouter{
		cardService,
	}
}

func (router *CardRouter) HandleRoutes(route *gin.RouterGroup) {
	cardApi := route.Group("/cards")
	{
		cardApi.GET("", router.getAll)
		cardApi.POST("", router.create)
	}
}

func (router *CardRouter) getAll(context *gin.Context) {
	cards, _ := router.cardService.GetAll()
	context.JSON(http.StatusOK, cards)
}

func (router *CardRouter) create(context *gin.Context) {
	var request domain.CardCreateReq
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, _ := router.cardService.Create(request)
	context.JSON(http.StatusCreated, card)
}
