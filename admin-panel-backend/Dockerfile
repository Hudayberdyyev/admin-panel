FROM golang:1.16-buster AS build

RUN go version

COPY . /github.com/Hudayberdyyev/admin-panel-backend/
WORKDIR /github.com/Hudayberdyyev/admin-panel-backend/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o backend cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=build /github.com/Hudayberdyyev/admin-panel-backend/backend .
COPY --from=build /github.com/Hudayberdyyev/admin-panel-backend/configs configs/
COPY --from=build /github.com/Hudayberdyyev/admin-panel-backend/logo logo/

CMD ["./backend"]