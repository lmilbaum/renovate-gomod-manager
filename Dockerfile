FROM golang:1.21.4@sha256:9baee0edab4139ae9b108fffabb8e2e98a67f0b259fd25283c2a084bd74fea0d

COPY main.go ./

CMD [ "go", "run", "./main.go", "/workspace" ]