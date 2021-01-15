FROM golang:1.15-alpine3.12

LABEL maintainer="onism68@gmail.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 获取变量
ARG socksProxy
ARG tgToken

ENV socksProxy=$socksProxy
ENV tgToken=$tgToken

# 设置固定的项目路径
ENV WORKDIR /telegram-bot

# 添加文件进项目路径
ADD . $WORKDIR

# 添加应用可执行文件，并设置执行权限
#ADD ./bin/linux_amd64/main   $WORKDIR/main
#RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
#ADD i18n     $WORKDIR/i18n
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD template $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
RUN go env -w GOPROXY=https://goproxy.cn,direct
ENTRYPOINT ["go", "run", "main.go"]
