svc.{{ .Client }}.{{ .Method }}({{- range .Params }}{{ . }}, {{- end }}nil)