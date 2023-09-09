FROM node:18


#  docker build --platform linux/amd64 . -t mojeico/nodejs-server:v666
#  docker push mojeico/nodejs-server:v666




WORKDIR /usr/src/app


COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3001
CMD [ "node", "server.js" ]