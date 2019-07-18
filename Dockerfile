FROM golang:1.12.5 as builder
ENV GO111MODULE=on
WORKDIR /go/src/my-go-project
ADD . /go/src/my-go-project
RUN make

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/my-go-project/my-go-project /app
COPY --from=builder /go/src/my-go-project/conf /app/conf/
EXPOSE 9990
ENTRYPOINT [ "./my-go-project", "-c", "/app/conf/config.prod.yaml" ]
