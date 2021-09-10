FROM golang:alpine

LABEL maintainer="Ex0dIa-dev"

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /app/ssh-honeypot-go .

ENTRYPOINT [ "/app/ssh-honeypot-go" ]
