FROM golang:1.23.1

WORKDIR /server

ENV PROJECT_DIR=/server \
    GO111MODULE=on \
    CGO_ENABLED=0

ENV GOFLAGS=-buildvcs=false

COPY . .

EXPOSE 3000

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

RUN go install github.com/a-h/templ/cmd/templ@latest

ENTRYPOINT CompileDaemon -include="*.templ" -exclude="*_templ.go" -build="templ generate && go build -o serve ." -command="./serve"
