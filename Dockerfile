FROM library/golang

RUN export GOPATH=$(go env GOPATH)
RUN export GOROOT=$(go env GOROOT)
ENV PATH=$PATH:$GOPATH/bin:$GOROOT/bin
ENV APP_DIR $GOPATH/src/sina/gobot
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./gobot)
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && go build -v .
