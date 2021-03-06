FROM golang:1.16
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ADD conf /go/conf
ADD app /app
RUN ["mkdir","/go/modules"]
EXPOSE 8138
CMD ["/app"]