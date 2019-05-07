# --cap-add NET_ADMIN
FROM library/golang:1.12

WORKDIR /go/src/github.com/cmattoon/conntrackr

COPY . .

RUN go build ./...
