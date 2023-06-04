package httplog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type HanlderOptions func(h *Handler)

func PathSize(size int) HanlderOptions {
	return func(h *Handler) {
		h.PathSize = size
	}
}

func NewHandler(handler http.Handler, options ...HanlderOptions) Handler {
	h := Handler{
		Handler:  handler,
		PathSize: 25,
	}

	for _, opt := range options {
		opt(&h)
	}

	return h
}

type Handler struct {
	Handler  http.Handler
	PathSize int
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rw := NewResponseWriter(w)

	h.Handler.ServeHTTP(rw, r)

	dur := time.Since(start)

	fmt.Println(
		"[httplog]",
		start.Format(time.DateTime),
		sprintMethod(r.Method),
		fmt.Sprintf("%-*s", h.PathSize, r.URL.Path),
		sprintStatus(rw.statusCode),
		dur,
	)
}

var METHOD_COLORS = map[string][]color.Attribute{
	http.MethodGet:    {color.BgGreen, color.FgBlack},
	http.MethodPost:   {color.BgYellow, color.FgBlack},
	http.MethodPut:    {color.BgBlue, color.FgBlack},
	http.MethodPatch:  {color.BgCyan, color.FgBlack},
	http.MethodDelete: {color.BgRed, color.FgBlack},
}

func sprintMethod(method string) string {
	c, ok := METHOD_COLORS[method]
	if !ok {
		c = []color.Attribute{color.BgWhite, color.FgBlack}
	}
	return color.New(c...).Sprintf(" %-7s ", method)
}

var STATUS_COLORS = map[int][]color.Attribute{
	2: {color.BgGreen, color.FgBlack},
	3: {color.BgBlue, color.FgBlack},
	4: {color.BgYellow, color.FgBlack},
	5: {color.BgRed, color.FgBlack},
}

func sprintStatus(code int) string {
	c := STATUS_COLORS[code/100]
	return color.New(c...).Sprintf(" %d ", code)
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, http.StatusOK}
}

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
