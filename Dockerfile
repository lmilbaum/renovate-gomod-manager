FROM golang:1.21.4@sha256:81cd210ae58a6529d832af2892db822b30d84f817a671b8e1c15cff0b271a3db

COPY main.go ./

CMD [ "go", "run", "./main.go", "/workspace" ]