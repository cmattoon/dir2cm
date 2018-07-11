FROM library/golang:1.9.2-alpine

LABEL org.label-schema.schema-version = "1.0.0"
LABEL org.label-schema.name = "dir2cm"
LABEL org.label-schema.description = "Creates a ConfigMap from a Directory"
LABEL org.label-schema.vendor = "com.cmattoon"
LABEL org.label-schema.vcs-url = "https://github.com/cmattoon/dir2cm"

WORKDIR /go/src/github.com/cmattoon/dir2cm

RUN apk add --no-cache git

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

VOLUME ["/data"]

WORKDIR /data

ENTRYPOINT ["dir2cm"]
