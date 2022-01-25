FROM golang:1.17

WORKDIR /
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN go install github.com/spf13/cobra/cobra@latest

CMD ["tail", "-f", "/dev/null"]