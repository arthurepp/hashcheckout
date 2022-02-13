FROM golang:1.17.7-alpine AS build-env
WORKDIR /app 
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o main . 
RUN CGO_ENABLED=0 GOOS=linux go test -mod=readonly -v

FROM alpine:3.15.0
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=build-env /app/main .
CMD ["./main"]