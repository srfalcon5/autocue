GOCMD= /opt/homebrew/bin/GOCMD

build:
	$(GOCMD) build -o server

clean:
	$(GOCMD) clean
	rm server