* Запуск без балансировщиком нагрузки
```bash
  wrk -c5 -t5 -d1m -s ./wrk.lua 'http://127.0.0.1:8081'
```
* результат 
```bash
Running 1m test @ http://127.0.0.1:8081
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.03ms    2.77ms 112.63ms   99.12%
    Req/Sec     1.11k   132.07     1.47k    83.61%
  331861 requests in 1.00m, 23.74MB read
Requests/sec:   5529.81
Transfer/sec:    405.02KB
```
* Запуск нескольких копий HTTP-сервера
```bash
docker-compose up -d --scale api=3
```
* Запуск с балансировщиком нагрузки
```bash
wrk -c5 -t5 -d1m -s ./wrk.lua 'http://127.0.0.1:4000'
```
* результат
```bash
Running 1m test @ http://127.0.0.1:4000
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.82ms    1.22ms  28.44ms   92.65%
    Req/Sec   587.13    131.24   767.00     78.33%
  175457 requests in 1.00m, 20.25MB read
Requests/sec:   2921.53
Transfer/sec:    345.21KB
```
* Вывод: из-за балансировщика нагрузки стало только хуже. Но вроде как значение Max снизилось.

* Запуск нескольких копий консюмера
* ВОПРОС: А это можно были ли как-то сделать через код? 
Я так понял чтобы несколько консюмеров в одной группе работали параллельно нужно создать несколько partitions.
Но вот как их создавать через код, так и не понял.
```bash
docker exec -it lesson7_kafka_1 /bin/bash
kafka-topics.sh --describe --zookeeper zookeeper:2181 --topic rates
kafka-topics.sh --list --zookeeper zookeeper:2181
kafka-topics.sh --zookeeper zookeeper:2181 --alter --topic rates --partitions 2 
docker-compose up -d --scale api=3 --scale process=2
```
