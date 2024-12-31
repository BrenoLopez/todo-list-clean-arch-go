FROM golang:1.23.4-alpine3.20 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /todo-application


FROM gcr.io/distroless/static-debian12 AS release-stage

WORKDIR /

COPY --from=build-stage /app/.env .env
COPY --from=build-stage /todo-application /todo-application

USER nonroot:nonroot

ENTRYPOINT ["/todo-application"]


FROM golang:1.23.4-alpine3.20 AS development

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

# ENV GO111MODULE=off
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o /todo-application .

CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/todo-application"]