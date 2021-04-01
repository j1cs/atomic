FROM golang:alpine
COPY . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 9001
CMD ["/app/main"]