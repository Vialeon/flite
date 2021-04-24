package ext

import (
	"strings"

	"github.com/augmentable-dev/flite/internal/http_request"
	"github.com/augmentable-dev/flite/internal/lines"
	"github.com/augmentable-dev/flite/internal/readfile"
	"github.com/augmentable-dev/flite/internal/str"
	"github.com/augmentable-dev/flite/internal/yaml"

	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func init() {
	split := func(s, c string, i int) string {
		split := strings.Split(s, c)
		if i < len(split) {
			return split[i]
		}
		return ""
	}
	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateModule("lines", lines.NewVTab(),
			sqlite.EponymousOnly(true), sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("readfile", readfile.NewReadFile()); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("yaml_to_json", yaml.NewYAMLToJSON()); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("json_to_yaml", yaml.NewJSONToYAML()); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		if err := api.CreateFunction("http_get", http_request.New_get()); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		if err := api.CreateFunction("http_post", http_request.New_post()); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		if err := api.CreateFunction("str_split", str.Newstr_split); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}
func main() {}
