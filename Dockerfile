FROM golang:1.21.4@sha256:57bf74a970b68b10fe005f17f550554406d9b696d10b29f1a4bdc8cae37fd063

COPY main.go ./

CMD [ "go", "run", "./main.go", "/workspace" ]