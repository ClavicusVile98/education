network


## Лабораторная работа #1
**Задание 1**

*Задание:*

- разобрать структуру приложенного Vagrantfile;
- нарисовать схему;
- расписать возможные подсети.

Схема сети:

![network.png](../_resources/02bad8790a1d4b85a3500b88b0bd6896.png)

**Задание 2**

- найти свободные подсети;
- подсчитать, сколько узлов в каждой подсети, включая свободные;
- указать broadcast-адрес для каждой подсети;
- проверить, нет ли ошибок при разбиении.

Подсети первой сети:

192.168.2.0/26, узлов: 62
broadcast: 192.168.2.63

192.168.2.64/26, узлов: 62
broadcast: 192.168.2.127

192.168.2.128/26, узлов: 62
broadcast: 192.168.2.191

Подсети второй сети:

192.168.1.0/25, узлов: 126
broadcast: 192.168.1.127

192.168.1.128/26, узлов: 62
broadcast: 192.168.1.191

192.168.1.192/26, узлов: 62
broadcast: 192.168.1.255

Подсети третьей сети:

192.168.0.0/28, узлов: 14
broadcast: 192.168.0.15

**Задание 3**

*Задание:*

- все серверы и роутеры должны ходить в Интернет черз inetRouter;
- все серверы должны видеть друг друга;
- у всех новых серверов отключить дефолт на NAT (eth0), который Vagrant поднимает для связи;
- в README приложить скриншоты tracepath и ip r.

office1Server -> office2Server
![office1_office2.png](../_resources/82b7d2f35704485f9da66d633952f5f1.png)
![office1_ip_r.png](../_resources/9b984b5d29e14079be7ed999321355d0.png)

office2Server -> office1Server
![office2_ip_r.png](../_resources/b4ae805943044897a6475c38bd9885ee.png)
![office2_office1.png](../_resources/a760f1498b74469789cde3641480192f.png)

centralServer -> office1Server && office2Server 
![centrealServer.png](../_resources/71f7267993af4078b6e762741fb519cf.png)
![tracepath_central.png](../_resources/12bbf909c9244ff0b1ff865af5969507.png)

**Задание 4**

*Задание:*

- поднять nginx на officе2Server
- запретить office1Server ходить на office2Server на 80й порт, все остальные должны работать
- запретить office1Server отвечать на пинг, всем кроме inetRouter, но office1Server должен пинговать всех остальных

![curl_office2.png](../_resources/cdc6b92ff01c473a90339b553d928259.png)


