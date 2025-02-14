package print

import (
	"text/tabwriter"
	"text/template"

	"github.com/replicatedhq/replicated/pkg/types"
)

var registriesTmplSrc = `PROVIDER	ENDPOINT	AUTHTYPE
{{ range . -}}
{{ .Provider }}	{{ .Endpoint }}	{{ .AuthType }}
{{ end }}`

var registriesTmpl = template.Must(template.New("registries").Funcs(funcs).Parse(registriesTmplSrc))

func Registries(w *tabwriter.Writer, registries []types.Registry) error {
	if err := registriesTmpl.Execute(w, registries); err != nil {
		return err
	}

	return w.Flush()
}

var registryLogsTmplSrc = `DATE	IMAGE	ACTION	STATUS	SUCCESS
{{ range . -}}
{{ .CreatedAt }}	{{ if not .Image }}{{ else }}{{ .Image }}{{ end }} 	{{ .Action }}	{{ if not .Status }}{{ else }}{{ .Status }}{{ end }}	{{ .Success }}
{{ end }}`

var registryLogsTmpl = template.Must(template.New("registryLogs").Funcs(funcs).Parse(registryLogsTmplSrc))

func RegistryLogs(w *tabwriter.Writer, logs []types.RegistryLog) error {
	if err := registryLogsTmpl.Execute(w, logs); err != nil {
		return err
	}

	return w.Flush()
}
