cd ./client 
yarn install 
cd ../server
cd glide install
cd ..
docker-compose -f docker-compose.dev.yml build

cd ./client && yarn start & docker-compose -f docker-compose.dev.yml up && fg