FROM golang:alpine as builder

LABEL maintainer="Ex0dIa-dev"

WORKDIR /build

#installing modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY src/ src/

#building
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags -s -w' -o /build/ssh-honeypot-go .

# second stage
FROM scratch

ENV HONEYPOT_CONFIGFILE="/app/config/config.json"
ENV HONEYPOT_HOSTKEYFILE="/app/config/hostkey_rsa"
ENV HONEYPOT_LOGSPATH="/app/logs/"
ENV HONEYPOT_CMDFILE="/app/config/cmds.txt"

COPY --from=builder /build/ssh-honeypot-go /app/

ENTRYPOINT [ "/app/ssh-honeypot-go" ]
