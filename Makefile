GOCMD= $(shell which go)
BROTLICMD= $(shell which brotli)

build:
	$(GOCMD) build main.go -o server
	$(GOCMD) build cleanup.go -o cleanup
	find web/ -type f -regex '.*\.\(html\|txt\|js\|css\|)$' -exec $(BROTLI) -fkv {} \;
	find web/app.css - type f -exec $(BROTLI) -fkv {} \;

clean:
	$(GOCMD) clean
	rm server
	rm cleanup
