# Часть 1. Round-Robin Balancer

Тестовое задание на стажировку в cloud.ru. Round-Robin балансировщик с Health Checks для проверки доступности серверов

### Как запустить и отправить запрос
- `docker-compose up --build`
- `curl http://localhost:8080`

### Что может пригодиться
- `docker stop test-service1` - остановить контейнер, т.е. сделать его недоступным
- `docker start test-service1` - запустить контейнер
- `docker kill --signal HUP balancer` - отправить сигнал SIGHUP для перезагрузки конфига 

Сейчас в конфиге 5 работающих серверов и 2 несуществуещих сервера

### Benchmark

This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 500 requests
Completed 1000 requests
Completed 1500 requests
Completed 2000 requests
Completed 2500 requests
Completed 3000 requests
Completed 3500 requests
Completed 4000 requests
Completed 4500 requests
Completed 5000 requests
Finished 5000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        0 bytes

Concurrency Level:      1000
Time taken for tests:   5.033 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      375000 bytes
HTML transferred:       0 bytes
Requests per second:    993.41 [#/sec] (mean)
Time per request:       1006.631 [ms] (mean)
Time per request:       1.007 [ms] (mean, across all concurrent requests)
Transfer rate:          72.76 [Kbytes/sec] received

Connection Times (ms)
min  mean[+/-sd] median   max
Connect:        0    7  14.6      0      62
Processing:    64  947 496.3    819    3081
Waiting:        2  946 496.1    818    3081
Total:         64  955 499.4    826    3100

Percentage of the requests served within a certain time (ms)
50%    826
66%    985
75%   1079
80%   1203
90%   1654
95%   1988
98%   2504
99%   2723
100%   3100 (longest request)
