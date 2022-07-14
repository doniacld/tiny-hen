#FROM golang:1.18
#
#RUN mkdir /app
#WORKDIR /app
#COPY tiny-hen tiny-hen
#EXPOSE 10010
#CMD ["tiny-hen"]

FROM scratch
ENTRYPOINT ["/tiny-hen"]
COPY tiny-hen /