FROM golang:1.18-alpine

WORKDIR /app
ADD /. /app

RUN go install

# swagger generator

RUN go build -o index

CMD ["/app/index"]
EXPOSE 8400