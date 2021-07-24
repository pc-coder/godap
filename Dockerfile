FROM golang:1.16-alpine3.14 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -o /out/mock-server .

FROM scratch AS bin

COPY --from=build /out/mock-server /

ENTRYPOINT ["/mock-server"]