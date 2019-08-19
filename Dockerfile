FROM golang:1.12.8-alpine3.10
LABEL maintainer="wormazabal <www.zeke.cl>"
LABEL version="1.0"
LABEL project="Web-GO-Backend-Test"
LABEL description="Imagen Docker para pruebas Go"
RUN mkdir /app
ADD . /app
WORKDIR /app
EXPOSE 3000/tcp
RUN apk add curl git vim wget
RUN go get -u github.com/go-chi/chi
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/swaggo/http-swagger
RUN go get -u github.com/alecthomas/template
RUN go get -u github.com/go-chi/cors
RUN go build -o main .
CMD ["/app/main"]