GOCMD= $(shell which go) 

build:
	$(GOCMD) build main.go -o server
	$(GOCMD) build cleanup.go -o cleanup

clean:
	$(GOCMD) clean
	rm server
	rm cleanup
