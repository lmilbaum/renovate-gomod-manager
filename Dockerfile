FROM golang:1.21.3@sha256:24a09375a6216764a3eda6a25490a88ac178b5fcb9511d59d0da5ebf9e496474

COPY main.go ./

CMD [ "go", "run", "./main.go" ]