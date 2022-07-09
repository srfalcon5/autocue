GOCMD= /opt/homebrew/bin/go

build:
	$(GOCMD) build -o server

clean:
	$(GOCMD) clean
	rm server