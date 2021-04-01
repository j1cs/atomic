FROM php:alpine
COPY . /app
WORKDIR /app
EXPOSE 9004
ENTRYPOINT ["php"]
CMD ["-S", "0.0.0.0:9004", "-t" ,"."]