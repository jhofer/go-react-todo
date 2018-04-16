cwd=$(PWD)
docker-compose -f docker-compose.dev.yml build
cd ./client && yarn install 
yarn install 
cd $cwd
cd ./client && yarn start & docker-compose -f docker-compose.dev.yml up && fg