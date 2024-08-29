FROM scratch
COPY sttr /usr/bin/sttr
ENTRYPOINT ["/usr/bin/sttr"]