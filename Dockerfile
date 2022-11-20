ARG APPNAME=app
# builder image
FROM golang:1.18-alpine3.15 as builder
ARG APPNAME
ENV GO111MODULE=on
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ${APPNAME} main.go


# generate clean, final image for end users
FROM alpine:3.11.3
ARG APPNAME

RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata



RUN mkdir /app
ADD resources/db/migrations /app/resources/db/migrations
ADD config /app/config

WORKDIR /app/
COPY entrypoint.sh /app/
RUN chmod +x entrypoint.sh
COPY --from=builder /build/${APPNAME} .

RUN rm -rf /var/cache/apk/*
# executable
EXPOSE 8003
ENTRYPOINT [ "./entrypoint.sh" ]
# arguments that can be overridden
CMD [ "serve" ]