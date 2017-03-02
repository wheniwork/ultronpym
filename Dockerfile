FROM golang:1.7
MAINTAINER Daniel Olfelt <dolfelt@gmail.com>

ENV APP_PATH=/go/src/github.com/wheniwork/ultronpym

WORKDIR $APP_PATH
COPY . $APP_PATH
RUN glide install

CMD ["go", "build"]
