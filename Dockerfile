FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /sort

WORKDIR /sort

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/sort


FROM scratch

COPY --from=builder /go/bin/sort /go/bin/sort

ENTRYPOINT ["/go/bin/sort"]

EXPOSE 8080
