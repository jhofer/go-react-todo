# Stage 0, based on Node.js, to build and compile the client
FROM node:9.11-alpine as node

RUN apk update && \
  apk upgrade && \
  apk add git g++ gcc libgcc libstdc++ linux-headers make python

COPY . /app
WORKDIR /app


RUN yarn install
#ARG env=prod
RUN yarn dist

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:alpine
RUN apk add --no-cache bash
COPY --from=node /app/dist/ /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./wait-for-it.sh ./wait-for-it.sh
RUN chmod +x ./wait-for-it.sh
CMD ["/wait-for-it.sh", "golang:3001", "--", "nginx", "-g", "daemon off;"]