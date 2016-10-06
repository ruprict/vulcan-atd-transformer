package transform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/vulcand/oxy/utils"
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

type bufferWriter struct {
	header http.Header
	code   int
	buffer *bytes.Buffer
}

func (b *bufferWriter) Close() error {
	return nil
}

func (b *bufferWriter) Header() http.Header {
	return b.header
}

func (b *bufferWriter) Write(buf []byte) (int, error) {
	return b.buffer.Write(buf)
}

// WriteHeader sets rw.Code.
func (b *bufferWriter) WriteHeader(code int) {
	b.code = code
}

func (h *TransformHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(InventoryResponse{4, time.Now()})
	fmt.Println("*** atd_transformer middleware ***")

	bw := &bufferWriter{header: make(http.Header), buffer: &bytes.Buffer{}}
	newBody := bytes.NewBufferString("")
	if err := applyString(string(json), newBody, r); err != nil {
		fmt.Errorf("can't write bddy")
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(newBody.Len()))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	utils.CopyHeaders(w.Header(), bw.Header())
	io.Copy(w, newBody)
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

func applyString(in string, out io.Writer, request *http.Request) error {
	t, err := template.New("t").Parse(in)
	if err != nil {
		return err
	}

	if err = t.Execute(out, data{request}); err != nil {
		return err
	}

	return nil
}

// data represents template data that is available to use in templates.
type data struct {
	Request *http.Request
}
