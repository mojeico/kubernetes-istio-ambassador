FROM node:18


#  docker build --platform linux/amd64 . -t mojeico/nodejs-client:v1
#  docker push mojeico/nodejs-client:v1



WORKDIR /usr/src/app


COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000
CMD [ "node", "server.ts" ]