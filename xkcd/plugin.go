package xkcd

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-xkcd",
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"xkcd_comic": tableXkcdComic(ctx),
		},
	}
	return p
}
