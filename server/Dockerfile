# the alpine image are minimal images based on alpine linux with a complete package index and only 5 MB in size!
FROM golang:alpine
# install git
RUN apk add --update tzdata \
  bash wget curl git;

# Create binary directory, install glide & gin(hotreload via gin-tonic/gin)
RUN mkdir -p $$GOPATH/bin && \
  curl https://glide.sh/get | sh && \
  go get github.com/codegangsta/gin

# define work directory
ADD . /go/src/app
WORKDIR /go/src/app

#install depdencies
RUN glide update 
#Compile app
RUN go install
CMD $GOPATH/bin/app
