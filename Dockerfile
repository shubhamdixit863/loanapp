# Building the binary of the App
FROM golang:1.18 AS build



# `boilerplate` should be replaced with your project name
WORKDIR /src

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

WORKDIR /src/cmd/web

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM ubuntu:latest

RUN apt-get update


RUN apt-get -y install openssl


RUN openssl s_client -showcerts -connect registry-1.docker.io:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >ca.crt
RUN cp ca.crt /etc/ssl/certs/ca.crt
RUN cat ca.crt |  tee -a /etc/ssl/certs/ca-certificates.crt

WORKDIR /app

# `boilerplate` should be replaced here as well
COPY --from=build /src/cmd/web/ .

# Exposes port 8090 because our program listens on that port
EXPOSE 8090

CMD ["./app"]