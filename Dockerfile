FROM golang:1.16
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ADD app /app
EXPOSE 8080
CMD ["/app"]