FROM golang:1.13

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -o myapp2

FROM alpine

COPY --from=0 /src/myapp2 /bin/myapp2
EXPOSE 5000:5000
ENTRYPOINT [ "/bin/myapp2" ]

