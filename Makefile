.PHONY: run
run:
	sudo service postgresql start
	go build && ./hori-api