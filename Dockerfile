FROM busybox
COPY ./ddtalk /apps
CMD ["/ddtalk"]