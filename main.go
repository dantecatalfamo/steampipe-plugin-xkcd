package main

import (
	"github.com/dantecatalfamo/steampipe-plugin-xkcd/xkcd"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: xkcd.Plugin})
}
