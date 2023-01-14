
# Kafka Simple

Example Producer parameter

```http
  POST /send-message
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `message`    | `string` | **Required**. Message    |
| `topic`    | `json` | **Required**. simple / simple-2

## Running 

To run tests, run the following command

```bash
  make run-consumer
```

```bash
  make run-producer
```

