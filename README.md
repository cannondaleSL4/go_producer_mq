#Create proto

protoc --go_out=paths=source_relative:. **/*.proto


docker run --rm -it -p 15672:15672 -p 5672:5672 rabbitmq:3-management
