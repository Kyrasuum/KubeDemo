# arguements
ARG IMAGE_NAME=alpine
ARG TAG=latest
# Name the node stage "builder"
FROM ${IMAGE_NAME}:${TAG} AS build-stage

RUN apk add nodejs npm git go --no-cache
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

WORKDIR /app/internal/pages
COPY package*.json ./
RUN npm install
RUN npm run build

WORKDIR /app
RUN go build -o web cmd/main.go

# production stage
FROM ${IMAGE_NAME}:${TAG} AS production-stage
WORKDIR /app
COPY --from=build-stage /app/web /app/
RUN adduser -D golanguser &&\
    chown -R golanguser /app
USER golanguser
EXPOSE 8081
CMD ["/app/web"]
