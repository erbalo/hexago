package card

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	lambdaAdapter "github.com/erbalo/hexago/internal/adapters/lambda"
	"github.com/erbalo/hexago/internal/app/card"
)

type Handler struct {
	service card.Service
}

func NewGetAllHandler(service card.Service) lambdaAdapter.Entrypoint {
	return &Handler{service}
}

func (handler *Handler) Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{}, nil
}

func (handler *Handler) Start() {
	lambda.Start(handler.Handler)
}
