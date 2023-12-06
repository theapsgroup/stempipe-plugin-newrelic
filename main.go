package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-newrelic/newrelic"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: newrelic.Plugin})
}
