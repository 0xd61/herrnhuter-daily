FROM golang:alpine as build
ARG VERSION=latest
RUN apk add --update git
WORKDIR /go/src/github.com/Kaitsh/herrnhuter-daily
COPY .    .
RUN go get -v -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-X main.version=$VERSION" -a -installsuffix cgo -o server-docker-v$VERSION .

FROM scratch
LABEL Maintainer=daniel.glinka@daimler.com
LABEL OWNER=DGLINKA
ARG VERSION=latest
COPY --from=build /go/src/github.com/Kaitsh/herrnhuter-daily/server-docker-v$VERSION /run/server
WORKDIR /run

CMD [ "./server" ]

