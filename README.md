# kafka-fake-message
Kafka Fake Message Generator For Test

Example1:

```bash
go run main.go start --brokers 127.0.0.1:9092 --topic example-topic --interval 2 --count 5 --file example.log
```

Example2:

```bash
go run main.go start --brokers 127.0.0.1:9092 --topic example-topic --interval 2 --count 5 --generate
```

Example3:

```bash
go run main.go start --brokers 127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094 --topic example-topic --interval 2 --count 5 --file example.log
```