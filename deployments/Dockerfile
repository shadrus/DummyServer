FROM golang:1.18-stretch as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./src/cmd/dummyserver

FROM alpine AS bin
COPY --from=build /out/app .
CMD ["./app"]