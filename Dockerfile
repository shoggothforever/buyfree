FROM golang:1.20
WORKDIR usr/src/app
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GO_PROXY=https://proxy.golang.com.cn,direct
ENV GO111MODULE=auto

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /bf

CMD ["/bf"]