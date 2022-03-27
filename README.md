## Notes and Contributing

---

Notes and contribution convention.

### **MQTT**

- Topic

  Mqtt topic uses slash as separator. e.g.

  - `air/humid` stands for air humid
  - `air/temp` stands for air temperature

### **AMQP**

- Queue

  In a nutshell Queue is another term of topic but in AMQP. Since we're using RabbitMQ as the broker it will automatically convert mqtt topic into amqp queue convention. Queue convention looks like these:

  - `air.humid` stands for air humid
  - `air.temp` stands for air temperature

  It uses dot instead of slash

### Notes

- Always use `.` as separator when creating queue that will be used from mqtt. Since it will be processed automatically by RabbitMQ.
