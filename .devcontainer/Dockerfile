FROM golang:1.20

RUN go install -v github.com/cweill/gotests/gotests@latest
RUN go install -v github.com/fatih/gomodifytags@latest
RUN go install -v github.com/josharian/impl@latest
RUN go install -v github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest 
RUN go install -v golang.org/x/tools/gopls@latest 
RUN go install -v github.com/ramya-rao-a/go-outline@latest