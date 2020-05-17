FROM golang:1.13

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -o myapp

FROM alpine

COPY --from=0 /src/myapp /bin/myapp
EXPOSE 4000:4000
ENTRYPOINT [ "/bin/myapp" ]

