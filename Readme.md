# Демо по Docker

## ENTRYPOINT vs CMD
1. Соберем образ cmd, для этого в директории cmd выполним команду `docker build -t cmd .`
2. Запустим контейнер `docker run cmd`, затем запустим с переопределением команды запуска `docker run cmd ping www.google.com`
3. Заменим содержимое Dockerfile на \
`FROM alpine:3.23` \
`ENTRYPOINT ["ping"]` \
`CMD ["www.ya.ru"]`
4. Пересоберем образ docker `build -t cmd .`
5. Запустим контейнер `docker run cmd`, затем запустим с переопределением дефолтного параметра `docker run cmd www.google.com -s 1`

## Сборка hello-world
1. Обратим внимание на последовательность команд в Dockerfile
2. Соберем образ `docker build -t hello-world:1.0 .`, посмотрим размер образа `docker images`
3. Соберем образ с использованием multi-stage билда `docker build -t hello-world:2.0 -f multi-stage/Dockerfile .`
4. Сравним получившийся размер образа с предыдущей версией `docker images`

## Запуск Nginx
1. Запустим контейнер с Nginx `docker run -d -p 8000:80 --name web-server nginx`, убедимся, что сервер отвечает `curl localhost:8000`
2. Посмотрим, как работают команды `docker ps`, `docker stats web-server`, `docker logs web-server`, `docker inspect web-server`
3. Подключимся к контейнеру `docker exec -it web-server sh`, убедимся, что находимся внутри контейнера `cat /etc/os-release`
4. Установим пакет procps (утилита ps) `apt-get update && apt-get install procps`, посмотрим запущенные процессы `ps -ef`
5. Посмотрим настройки Nginx, для этого перейдем в директорию `cd /etc/nginx` и выведем на экран конфиг `cat nginx.conf`
6. Посмотрим, как настроено логирование Nginx, для этого перейдем в директорию `cd /var/log/nginx` и выполним `ls -l`
7. Заменим index.html, который использует Nginx:
   1. Nginx берет index.html из директории `/usr/share/nginx/html`
   2. Остановим контейнер с Nginx `docker stop web-server`, затем удалим `docker rm web-server`
   3. Запустим Nginx с монтированием каталога с альтернативным index.html `docker run -d -p 8000:80 -v ./nginx-html:/usr/share/nginx/html --name web-server nginx`
   4. Убедимся, что Nginx возвращает новую html страницу при открытии `curl localhost:8000`
   5. `docker stop web-server && docker rm web-server`

## Ограничение ресурсов
1. Запустим hello-world `docker run -d --name hello-world hello-world:2.0`, посмотрим, как увеличивается потребление памяти `docker stats hello-world`
2. Остановим и удалим hello-world `docker stop hello-world && docker rm hello-world`
3. Запустим hello-world с ограничением используемой контейнером памяти `docker run -d -m 256m --name hello-world hello-world:2.0`
4. Посмотрим потребление ресурсов `docker stats hello-world`, убедимся, что контейнер использует не более 256 Мб памяти

## Docker сети
1. Создадим сеть `docker network create --driver=bridge my-net`
2. Запустим Nginx с указанием сети `docker run -d --network=my-net --name web-server nginx`
3. Запустим еще один контейнер в той же сети в интерактивном режиме `docker run -it --network=my-net alpine/curl sh`
4. Убедимся, что можем достучаться до контейнера с Nginx по его имени `curl web-server:80`
5. Остановим и удалим Nginx `docker stop web-server && docker rm web-server`
6. Запустим Nginx без указания сети `docker run -d --name web-server nginx`, и найдем его ip адрес `docker inspect web-server | grep IPAddress`
7. Запустим второй контейнер `docker run -it alpine/curl sh`, убедимся, что не можем достучаться до Nginx по имени `curl web-server:80`, а по ip адресу можем `curl <ip>:80`
8. Остановим и удалим Nginx `docker stop web-server && docker rm web-server`

## Docker compose
1. Посмотрим файл `docker-compose/compose.yaml`
2. Запустим все сервисы `docker-compose up -d`
3. Посмотрим логи всех запущенных сервисов `docker-compose logs`, проверим запущенные контейнеры `docker ps` и существующие сети `docker network ls`
4. Остановим и удалим все сервисы `docker-compose down`