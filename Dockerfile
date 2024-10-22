FROM golang:1.23.0-alpine

WORKDIR /app

COPY . .

RUN go build -o /hello-world-app

CMD [ "/hello-world-app" ]