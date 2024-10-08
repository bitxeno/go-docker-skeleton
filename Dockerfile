# TODO: 支持arm64: https://github.com/sgerrand/alpine-pkg-glibc/issues/126
FROM woahbase/alpine-glibc:2.39
ARG APP_NAME
ARG VERSION
ARG BUILDDATE
ARG COMMIT
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
RUN echo "I'm building for $TARGETPLATFORM"

# 安装tzdata支持更新时区
RUN apk add -U tzdata

# add 指令会自动解压文件
COPY ./docs/config.yaml.example /docs/config.yaml
COPY ./build/${APP_NAME}-${TARGETOS}-${TARGETARCH} /usr/bin/${APP_NAME}
RUN chmod +x /usr/bin/${APP_NAME}

# 生成启动脚本
RUN printf '#!/bin/sh \n\n\

if [ ! -f "/data/config.yaml" ]; then  \n\
    cp /docs/config.yaml /data/config.yaml \n\
fi  \n\

/usr/bin/%s server -p ${SERVICE_PORT:-80} -c /data/config.yaml  \n\
\n\
' ${APP_NAME} >> /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
# docker 启动不了，需要进入 docker 测试时使用本命令
# docker run -it --entrypoint /bin/sh [docker_image]

ENV PUID=0 PGID=0 UMASK=022
EXPOSE 80
VOLUME /data