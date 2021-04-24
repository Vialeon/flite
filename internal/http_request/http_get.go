package http_request

import (
	"io/ioutil"
	"net/http"
	"os"

	"go.riyazali.net/sqlite"
)

type http_get struct{}

// TODO add PUT and POST stuff

func (m *http_get) Args() int           { return -1 }
func (m *http_get) Deterministic() bool { return false }
func (m *http_get) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		request  string
		err      error
		contents []byte
		response *http.Response
	)

	if len(values) > 0 {
		request = values[0].Text()
	}
	print("and the request is \n", request)
	if request == "" {
		contents, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			ctx.ResultError(err)
		}
		response, err = http.Get(string(contents))

	} else {
		response, err = http.Get(request)
		if err != nil {
			println("there was an error", err)
			ctx.ResultError(err)
		}
		contents, err = ioutil.ReadAll(response.Body)
		if err != nil {
			ctx.ResultError(err)
		}
		println(string(contents))

	}

	ctx.ResultText(string(contents))
}

// Newhttp_get returns a sqlite function for reading the contents of a file
func New_get() sqlite.Function {
	return &http_get{}
}
