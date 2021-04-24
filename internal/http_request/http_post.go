package http_request

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"go.riyazali.net/sqlite"
)

type http_post struct{}

// TODO add PUT and POST stuff

func (m *http_post) Args() int           { return -1 }
func (m *http_post) Deterministic() bool { return false }
func (m *http_post) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		request      io.Reader
		url          string
		err          error
		content_type string
		contents     []byte
		response     *http.Response
	)
	for _, v := range values {
		print(v.Text(), "\n")
	}
	if len(values) > 2 {
		url = values[0].Text()
		content_type = values[1].Text()
		request = strings.NewReader(values[2].Text())
	}
	print("url is ", url, " content type is ", content_type, " request is ", request)
	print("and the request is \n", request)
	response, err = http.Post(url, string(content_type), request)
	if err != nil {
		contents, _ = ioutil.ReadAll(response.Body)
		println("Post request returned an error", err.Error(), "with code ", contents)
		ctx.ResultError(err)
	}
	contents, err = ioutil.ReadAll(response.Body)
	if err != nil {
		ctx.ResultError(err)
	}

	ctx.ResultText(string(contents))
}

// Newhttp_post returns a sqlite function for reading the contents of a file
func New_post() sqlite.Function {
	return &http_post{}
}
