FROM golang:latest
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .cmd/main
CMD .\out\dist


#RUN apk --no-cache add gcc g++ make git
#WORKDIR /go/src/app
#COPY . .
#RUN go mod init webserver
#RUN go mod tidy
#RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go
#
#FROM alpine:3.17
#RUN apk --no-cache add ca-certificates
#WORKDIR /usr/bin
#COPY --from=build /go/src/app/bin /go/bin
#EXPOSE 80
#ENTRYPOINT /go/bin/web-app --port 80