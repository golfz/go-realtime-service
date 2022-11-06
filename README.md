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

## 2. การรัน Service

### (2.1) การกำหนด Environment Variables

- `REALTIME_AMQP_ENDPOINT` เช่น
    - `amqp://guest:guest@localhost:5672/realtime`
    - คือค่า AMQP URL เพื่อใช้ในการ connection
- `REALTIME_PORT` เช่น
    - `80`
    - คือค่า Port ที่จะใช้ในการรัน Service

### (2.2) ตัวอย่างการ Run เพื่อทดสอบ

```shell
REALTIME_AMQP_ENDPOINT="amqp://guest:guest@localhost:5672" REALTIME_PORT="80" go run application.go
```