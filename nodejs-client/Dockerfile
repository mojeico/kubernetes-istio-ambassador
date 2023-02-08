FROM node:18


#  docker build --platform linux/amd64 . -t europe-west3-docker.pkg.dev/mojeico/client/nodejs-client:v8
#  docker push europe-west3-docker.pkg.dev/mojeico/client/nodejs-client:v8





WORKDIR /usr/src/app


COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000
CMD [ "node", "server.ts" ]