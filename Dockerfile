#reference: https://cloud.docker.com/repository/registry-1.docker.io/wormaza/go-chi-swagger
FROM wormaza/go-chi-swagger:dev-1.0
LABEL maintainer="wormaza"
LABEL version="1.0"
LABEL project="Web-GO-Chi-Backend-Test"
LABEL description="Imagen Docker para pruebas Go"
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY /SSL /app/ssl
EXPOSE 443/tcp
RUN go build -o main .
CMD ["/app/main"]