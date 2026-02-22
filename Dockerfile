FROM golang:1.25 AS build

WORKDIR /app

COPY . .

RUN make all

FROM scratch

COPY --from=build /app/bin/gostub /gostub
COPY --from=build /app/bin/config.yaml /config.yaml
COPY --from=build /app/bin/stubs /stubs

CMD ["-config", "/config.yaml"]
ENTRYPOINT ["/gostub"]
