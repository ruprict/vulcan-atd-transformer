package transform

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/codegangsta/cli"
	"github.com/vulcand/vulcand/plugin"
)

const Type = "atd_transformer"

func GetSpec() *plugin.MiddlewareSpec {
	return &plugin.MiddlewareSpec{
		Type:      Type,
		FromOther: FromOther,
		FromCli:   FromCli,
		CliFlags:  CliFlags(),
	}
}

type TransformMiddleware struct {
}

type TransformHandler struct {
	config TransformMiddleware
	next   http.Handler
}

type InventoryResponse struct {
	Quantity              int
	EstimatedDeliveryDate time.Time
}

func (h *TransformHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(InventoryResponse{4, time.Now()})
	w.Write(json)
	return
}

func New() (*TransformMiddleware, error) {
	return &TransformMiddleware{}, nil
}

func (m *TransformMiddleware) NewHandler(next http.Handler) (http.Handler, error) {
	return &TransformHandler{next: next, config: *m}, nil
}

func FromOther(m TransformMiddleware) (plugin.Middleware, error) {
	return New()
}

func FromCli(c *cli.Context) (plugin.Middleware, error) {
	return New()
}

func CliFlags() []cli.Flag {
	return []cli.Flag{}
}
