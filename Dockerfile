FROM golang:1.21.3

COPY main.go ./

CMD [ "go", "run", "./main.go" ]