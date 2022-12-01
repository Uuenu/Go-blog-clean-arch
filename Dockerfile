FROM golang:1.19

WORKDIR /go-blog-ca

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /go-blog-ca/internal/app ./...

CMD ["app"]