cd ./client && yarn install && yarn start & docker-compose -f docker-compose.dev.yml build
 && docker-compose -f docker-compose.dev.yml up && fg