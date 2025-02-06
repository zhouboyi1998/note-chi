FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-chi
COPY ./main /go/note-chi
COPY ./application-docker.yaml /go/note-chi

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18082
ENTRYPOINT ["./main"]
