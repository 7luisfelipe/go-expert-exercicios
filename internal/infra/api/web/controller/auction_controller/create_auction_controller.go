package auction_controller

import (
	"context"

	"net/http"

	"github.com/7luisfelipe/projetoleilao/configuration/rest_err"
	"github.com/7luisfelipe/projetoleilao/internal/infra/api/web/validation"
	"github.com/7luisfelipe/projetoleilao/internal/usecase/auction_usecase"
	"github.com/gin-gonic/gin"
)

type AuctionController struct {
	auctionUseCase auction_usecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (u *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	auctionId, err := u.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	//c.Status(http.StatusCreated)
	c.JSON(http.StatusCreated, auctionId)
}
