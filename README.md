# Golang Realtime Service (Example)

โปรเจคนี้ใช้ Library จาก [github.com/golfz/gorealtime](https://github.com/golfz/gorealtime)

## 1. การรัน RabbitMQ ด้วย Docker

### Windows
```shell
docker run -d --restart always --name rabbitmq --hostname docker-rabbitmq -p 5672:5672 -p 15672:15672 -v d:/development/rabbitmq/docker/data:/var/lib/rabbitmq/mnesia rabbitmq:management
```
หมายเหตุ: สามารถเปลี่ยน `d:/development/rabbitmq/docker/data` เป็น path ที่ต้องการได้

### Linux และ Mac
```shell
docker run -d --restart always --name rabbitmq --hostname docker-rabbitmq -p 5672:5672 -p 15672:15672 -v ~/workspace/development/docker-rabbitmq/data:/var/lib/rabbitmq/mnesia rabbitmq:management
```
หมายเหตุ: สามารถเปลี่ยน `~/workspace/development/docker-rabbitmq/data` เป็น path ที่ต้องการได้