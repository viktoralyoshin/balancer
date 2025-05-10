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


Server Software:\
Server Hostname:        localhost\
Server Port:            8080\

Document Path:          /\
Document Length:        0 bytes

Concurrency Level:      1000\
Time taken for tests:   4.913 seconds\
Complete requests:      5000\
Failed requests:        0\
Total transferred:      375000 bytes\
HTML transferred:       0 bytes\
Requests per second:    1017.78 [#/sec] (mean)\
Time per request:       982.533 [ms] (mean)\
Time per request:       0.983 [ms] (mean, across all concurrent requests)\
Transfer rate:          74.54 [Kbytes/sec] received

Connection Times (ms)
min  mean[+/-sd] median   max
Connect:        0    5  11.2      0     117\
Processing:    57  933 335.7    901    2106\
Waiting:        3  932 335.6    901    2106\
Total:         57  938 340.0    903    2123

Percentage of the requests served within a certain time (ms)
50%    903\
66%   1042\
75%   1129\
80%   1191\
90%   1426\
95%   1555\
98%   1766\
99%   1825\
100%   2123 (longest request)
