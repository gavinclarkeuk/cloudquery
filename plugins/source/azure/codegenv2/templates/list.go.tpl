// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.SubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

{{if not .SkipFetch}}
func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().{{.Service | ToCamel}}{{.SubService | ToCamel}}
  pager := svc.{{.ListFuncName}}(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.{{.OutputField}}
	}
	return nil
}
{{end}}