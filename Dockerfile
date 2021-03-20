FROM golang:alpine

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o webcalculator cmd/fictional-chainsaw/*

CMD ["/app/webcalculator"]