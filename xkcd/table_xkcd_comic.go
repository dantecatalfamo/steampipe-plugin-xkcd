package xkcd


import (
	"context"

	xkcdClient "github.com/nishanths/go-xkcd/v2"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableXkcdComic(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "xkcd_comic",
		Description: "XKCD Comic",
		List: &plugin.ListConfig{
			Hydrate: getXkcd,
			KeyColumns: plugin.SingleColumn("number"),
		},
		Columns: []*plugin.Column{
			{ Name: "alt", Type: proto.ColumnType_STRING, Description: "Alternate Text" },
			{ Name: "day", Type: proto.ColumnType_INT, Description: "Day Published" },
			{ Name: "image_url", Type: proto.ColumnType_STRING, Description: "Image URL" },
			{ Name: "url", Type: proto.ColumnType_STRING, Description: "Comic URL" },
			{ Name: "month", Type: proto.ColumnType_INT, Description: "Month Published" },
			{ Name: "news", Type: proto.ColumnType_STRING, Description: "News" },
			{ Name: "number", Type: proto.ColumnType_INT, Description: "Comic Number" },
			{ Name: "safeTitle", Type: proto.ColumnType_STRING, Description: "Safe Title" },
			{ Name: "title", Type: proto.ColumnType_STRING, Description: "Title" },
			{ Name: "transcript", Type: proto.ColumnType_STRING, Description: "Transcript" },
			{ Name: "year", Type: proto.ColumnType_INT, Description: "Year Published" },
		},
	}
}

func getXkcd(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client := xkcdClient.NewClient()

	quals := d.KeyColumnQuals
	number := quals["number"].GetInt64Value()

	comic, err := client.Get(ctx, int(number))
	if err != nil {
		plugin.Logger(ctx).Error("xkcd.getXkcd", "query_error", err)
	}
	d.StreamListItem(ctx, comic)

	return nil, nil
}
