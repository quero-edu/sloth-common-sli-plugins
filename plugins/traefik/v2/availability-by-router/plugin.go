package availability

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

const (
	// SLIPluginVersion is the version of the plugin spec.
	SLIPluginVersion = "prometheus/v1"
	// SLIPluginID is the registering ID of the plugin.
	SLIPluginID = "sloth-common/traefik/v2/availability-by-router"
)

var queryTpl = template.Must(template.New("").Option("missingkey=error").Parse(`
(
  sum(rate(traefik_router_requests_total{ {{.filter}}router=~"{{.routerName}}",code=~"(5..)" }[{{"{{.window}}"}}]))
  /
  (sum(rate(traefik_router_requests_total{ {{.filter}}router=~"{{.routerName}}" }[{{"{{.window}}"}}])) > 0)
) OR on() vector(0)
`))

// SLIPlugin will return a query that will return the availability error based on traefik V2 router metrics.
func SLIPlugin(ctx context.Context, meta, labels, options map[string]string) (string, error) {
	router, err := getRouterName(options)
	if err != nil {
		return "", fmt.Errorf("could not get router name: %w", err)
	}

	// Create query.
	var b bytes.Buffer
	data := map[string]string{
		"filter":     getFilter(options),
		"routerName": router,
	}
	err = queryTpl.Execute(&b, data)
	if err != nil {
		return "", fmt.Errorf("could not render query template: %w", err)
	}

	return b.String(), nil
}

func getFilter(options map[string]string) string {
	filter := options["filter"]
	filter = strings.Trim(filter, "{},")
	if filter != "" {
		filter += ","
	}

	return filter
}

func getRouterName(options map[string]string) (string, error) {
	router := options["router_name_regex"]
	router = strings.TrimSpace(router)

	if router == "" {
		return "", fmt.Errorf("router name is required")
	}

	_, err := regexp.Compile(router)
	if err != nil {
		return "", fmt.Errorf("invalid regex: %w", err)
	}

	return router, nil
}
