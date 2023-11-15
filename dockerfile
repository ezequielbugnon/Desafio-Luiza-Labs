FROM Golang:latest

WORKDIR /app

COPY . .

RUN go get
RUN go build -o bin .

EXPOSE 8080

ENTRYPOINT [ "/app/bin" ]