FROM golang
ADD . /root/PDF2Image

WORKDIR /root/PDF2Image

RUN apt-get update \
    && apt-get install -y libmagickwand-dev \
    && cp policy.xml /etc/ImageMagick-6/policy.xml \
    && go env GOPROXY=https://goproxy.io,direct \
    && go get github.com/gin-gonic/gin \
    && go get gopkg.in/gographics/imagick.v2/imagick \
    && go install

ENTRYPOINT pdf2image