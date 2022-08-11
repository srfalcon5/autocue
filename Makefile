GOCMD= $(shell which go) 

build:
	$(GOCMD) build -o server

clean:
	$(GOCMD) clean
	rm server
