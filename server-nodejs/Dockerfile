FROM node:18


#  docker build --platform linux/amd64 . -t mojeico/nodejs-server:v1
#  docker push mojeico/nodejs-server:v1




WORKDIR /usr/src/app


COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3001
CMD [ "node", "server.js" ]