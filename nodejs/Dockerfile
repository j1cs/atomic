FROM node:alpine
COPY . /app
WORKDIR /app
RUN npm install
EXPOSE 9003
ENTRYPOINT [ "node" ]
CMD [ "index.js" ]