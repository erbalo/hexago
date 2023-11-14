package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Entrypoint interface {
	Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)
	Start()
}
