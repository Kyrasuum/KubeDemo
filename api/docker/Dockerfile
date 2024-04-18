# arguements
ARG IMAGE_NAME=alpine
ARG TAG=latest
# Name the golang builder stage
FROM ${IMAGE_NAME}:${TAG} AS build-stage

RUN apk add --no-cache git musl-dev bash gcc go
ENV GOPATH /go
ENV GOROOT /usr/lib/go
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /app
COPY go.mod .

COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN go build -o api cmd/main.go

# production stage
FROM ${IMAGE_NAME}:${TAG} AS production-stage
WORKDIR /app
COPY --from=build-stage /app/api /app/
COPY internal/data /app/data
RUN adduser -D golanguser &&\
    chown -R golanguser /app
USER golanguser
EXPOSE 8082
CMD ["/app/api"]
