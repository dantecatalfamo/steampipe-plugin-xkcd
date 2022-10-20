steampipe_dir := ~/.steampipe
install_dir := $(steampipe_dir)/plugins/local/github.com/dantecatalfamo/xkcd/

steampipe-plugin-xkcd.plugin: main.go go.mod go.sum xkcd/plugin.go xkcd/table_xkcd_comic.go
	go build -o $@

install: steampipe-plugin-xkcd.plugin xkcd.spc
	mkdir -p $(install_dir)
	install -d $(install_dir)
	install $< $(install_dir)
	install xkcd.spc $(steampipe_dir)/config/

.PHONY: install
