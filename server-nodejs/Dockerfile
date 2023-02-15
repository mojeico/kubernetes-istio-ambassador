FROM node:12


#  docker build --platform linux/amd64 . -t europe-west3-docker.pkg.dev/mojeico/client/nodejs-server:v1
#  docker push europe-west3-docker.pkg.dev/mojeico/client/nodejs-server:v1

# docker build  . -t nodejs-server:v1




WORKDIR /usr/src/app


COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3001
CMD [ "node", "server.js" ]