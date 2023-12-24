FROM golang:latest AS build 

WORKDIR /build

COPY . /build

RUN cd /build ; make build

# stage 2 

FROM gcr.io/distroless/base-debian12

USER nonroot:nonroot

WORKDIR /build

COPY --from=build /build/bin/ /build

CMD ["/build/email"]
