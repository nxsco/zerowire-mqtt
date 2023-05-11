FROM golang:1.19 AS build

COPY ./ /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o zw ./cmd/zerowire-mqtt

FROM scratch AS final
COPY --from=build /app/zw /app/zw
CMD [ "/app/zw" ]