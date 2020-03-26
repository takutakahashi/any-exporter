# build stage
FROM golang AS build-env
ADD . /src
RUN cd /src && GO111MODULE=on go build -o c cmd/cmd.go

# serve stage
FROM ubuntu
WORKDIR /app
COPY --from=build-env /src/c /app/cmd
COPY ./example/ /app/example/
CMD ["./cmd", "--config", "./example/config.yaml"]
