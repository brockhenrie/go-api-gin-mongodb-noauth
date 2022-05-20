To Setup Project

Run this docker commnand to setup mongo database and persist data to mongodb/data

docker run --name mongo-db -d -vPath/to\data\db:/data/db  -e MONGO_INITDB_ROOT_USERNAME=dev -e MONGO_INITDB_ROOT_PASSWORD=Password123!! -p 27017:27017 mongo:latest 

I skipped the config mongod.conf because i was getting issues with the config file

in the root of the directory run 

RUN go build

RUN ./api

in ---- Postman or curl ----   add some users to db

Post("localhost:9090/v1/user/create")
{
  "name": "Person 1",
  "age": 21,
  "address": {
    "address1": "123 w 1st st",
    "address2": "",
    "city": "phoenix",
    "state": "Arizona",
    "zip": 85213
  }
}

base path "localhost:9090/v1/user"
    Get(".../get/:name")
    Get(".../getall") 
    Patch(".../update") 
    Post(".../create")
    Deltet(".../delete/:name")