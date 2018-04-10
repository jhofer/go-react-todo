# Stage 0, based on Node.js, to build and compile the client
FROM node:9.11 as node

RUN yarn install
#ARG env=prod
RUN yarn dist

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:1.13
COPY --from=node /dist/ /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf