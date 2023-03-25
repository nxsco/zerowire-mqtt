FROM golang:1.19 AS build

COPY ./ /app
WORKDIR /app

RUN go build -o zerowire-mqtt ./cmd/zerowire-mqtt

FROM scratch AS final
COPY --from=build /app/zw /app/zw
CMD [ "/app/zw" ]