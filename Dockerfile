FROM golang:1.21.3@sha256:b113af1e8b06f06a18ad41a6b331646dff587d7a4cf740f4852d16c49ed8ad73

COPY main.go ./

CMD [ "go", "run", "./main.go", "/workspace" ]