GOCMD= $(shell which go) 

build:
	$(GOCMD) build -o server

clean:
	$(GOCMD) clean
	rm server

serve:
	cp Caddyfile web/Caddyfile
	rsync -avrz web/ git@maatt.fr:cue.f5.maatt.fr
