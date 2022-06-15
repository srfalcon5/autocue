[Demo](https://f5-prompter.maatt.fr)

## Building
0. Install [Golang](https://golang.org/dl)
1. Clone this repository (`git clone https://github.com/doamatto/falcon5-teleprompter.git f5-prompter`)
2. Build the code (`go build -o server .`)
3. Run the compiled server file (`./server`)

## Deploying
0. Install [Docker](https://docker.com) or a compatible engine
1. Build the image for Docker (`docker build -t f5-prompter .`)
2. Run the built image (`docker run -d f5-prompter`)

## Acknoledgements
Created by [Matt Ronchetto](https://maatt.fr) and [Daniel Calderon](https://daan.ws) for the class of 2023 and beyond.