package str

import (
	"strings"

	"go.riyazali.net/sqlite"
)

type str_split struct{}

func (m *str_split) Args() int           { return 1 }
func (m *str_split) Deterministic() bool { return true }
func (m *str_split) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	input := values[0].Text()
	delim := values[1].Text()

	split := strings.Split(input, delim)
	ctx.ResultText(strings.Join(split, ","))

}

// Newstr_split returns a sqlite function for converting json to yaml
func Newstr_split() sqlite.Function {
	return &str_split{}
}
