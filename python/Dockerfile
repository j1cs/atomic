FROM python:alpine
RUN apk add --update \
    py-pip
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
EXPOSE 9005
ENTRYPOINT ["python"]
CMD ["app.py"]