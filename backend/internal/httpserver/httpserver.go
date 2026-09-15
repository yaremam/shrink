// Package httpserver builds the Shrink! backend's HTTP handler tree.
package httpserver

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type healthOutput struct {
	Body struct {
		Status string `json:"status" example:"ok"`
	}
}

// New builds the Shrink! backend's HTTP handler.
func New() http.Handler {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Shrink! API", "0.1.0"))

	huma.Register(api, huma.Operation{
		OperationID: "health",
		Method:      http.MethodGet,
		Path:        "/health",
		Summary:     "Health check",
	}, func(ctx context.Context, input *struct{}) (*healthOutput, error) {
		resp := &healthOutput{}
		resp.Body.Status = "ok"
		return resp, nil
	})

	return mux
}
