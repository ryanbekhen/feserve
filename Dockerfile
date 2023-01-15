FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o feserve .

FROM alpine
COPY --from=build /app/feserve /usr/local/bin/feserve
EXPOSE 8000

ENTRYPOINT ["feserve"]
