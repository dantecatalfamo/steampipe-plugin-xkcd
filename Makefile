steampipe_dir := ~/.steampipe
install_dir := $(steampipe_dir)/plugins/local/github.com/dantecatalfamo/xkcd

steampipe-plugin-xkcd.plugin:
	go build -o $@

install: steampipe-plugin-xkcd.plugin xkcd.spc
	mkdir -p $(install_dir)
	cp $< $(install_dir)
	cp xkcd.spc $(steampipe_dir)/config/

.PHONY: install
