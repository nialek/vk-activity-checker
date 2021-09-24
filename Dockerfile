FROM alpine
EXPOSE 8080
ADD main /
CMD ["/main"]