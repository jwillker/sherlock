FROM alpine
COPY sherlock /sherlock
ENTRYPOINT [ "/sherlock" ]
