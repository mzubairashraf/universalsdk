# step 1: build
FROM golang:1.12-alpine3.10 as build-step

# for go mod download
RUN apk add --update --no-cache ca-certificates git
RUN mkdir /universalsdk

WORKDIR /universalsdk
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

#ENTRYPOINT go run main.go
#EXPOSE 8080

RUN CGO_ENABLED=0 go build -o /deploy/universalsdk
# -----------------------------------------------------------------------------
# step 2: exec
FROM scratch

COPY --from=build-step /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-step /deploy/universalsdk /deploy/universalsdk


ENTRYPOINT ["/deploy/universalsdk", "/deploy/"]
EXPOSE 8080