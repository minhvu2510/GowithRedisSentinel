Hướng dẫn cài đặt redis sentinel với docker theo mô hình + kết nối golang tới redis sentinel
+ 1 master
+ 2 slave
+ 3 redis sentinel



![image](https://user-images.githubusercontent.com/36092539/177714843-9c6de79c-6e4e-4ba6-853c-94b17792f645.png)


B1 : Clone project 

B2 : Chạy docker compose :

```
docker-compose up -d

```

Kiểm tra các container đã lên :

```
docker ps
```

![image](https://user-images.githubusercontent.com/36092539/177716763-690568f2-94f6-4785-9bf6-9433df093577.png)

Có 4 container trong đó :

1: redis_master1 : Cài đặt master và sentinel

2: redis_slave và redis_sentinel cài đặt slave và sentinel

3: core_go chạy code golang

B3: Cài đặt redis trên từng server:

1. Cài đặt node master

    * Truy cập vào container để cài đặt : ``` docker exec -it  redis_master1 bash ```
    * Cấu hình redis master bằng cách sửa thông tin file config như sau : vim /etc/redis/redis.conf
    * Các thông tin cần sửa: 
    ``` 
        # bind 127.0.0.1 ::1       # Để redis public ra ngoài cho slave có thể connect
        requirepass CkkUdLPMT         # Đặt password dùng cli
        masterauth CkkUdLPMT        # Đặt password cho node master

     ```
    * Cấu hình redis sentinel bằng cách sửa file config như sau : ``` vim /etc/redis/sentinel.conf ```
    ``` 
        sentinel monitor mymaster 172.31.0.2 6379 2    # 172.31.0.2 là ip node master lấy bằng lệnh ifconfig
        sentinel auth-pass mymaster CkkUdLPMT
        sentinel parallel-syncs mymaster 1

     ```
    * Start redis-server và redis-sentinel
     ``` 
        service redis-server start
        service redis-sentinel start
     ```
2. Cài đặt 2 node slave vào sentinel:
    * Truy cập vào 2 container redis_slave, redis_sentinel để cài đặt : ``` docker exec -it  redis_slave bash ```
    * Cấu hình redis master bằng cách sửa thông tin file config như sau : vim /etc/redis/redis.conf
    * Các thông tin cần sửa: 
    ``` 
        # bind 127.0.0.1 ::1       # Để redis public ra ngoài cho slave có thể connect
        requirepass CkkUdLPMT         # Đặt password dùng cli
        masterauth CkkUdLPMT        # Đặt password cho node master
        replicaof 172.31.0.2 6379   # 172.31.0.2 là ip node master

     ```
    * Cấu hình redis sentinel bằng cách sửa file config như sau : ``` vim /etc/redis/sentinel.conf ```
    ``` 
        sentinel monitor mymaster 172.31.0.2 6379 2    # 172.31.0.2 là ip node master lấy bằng lệnh ifconfig
        sentinel auth-pass mymaster CkkUdLPMT
        sentinel parallel-syncs mymaster 1

     ```
    * Start redis-server và redis-sentinel
     ``` 
        service redis-server start
        service redis-sentinel start
     ```
3. Kiểm tra cài đặt:
    * Truy cập vào container master : ``` docker exec -it  redis_master1 bash ```
    * Kiểm tra thông tin:

    ![image](https://user-images.githubusercontent.com/36092539/177722782-f1b62a65-8549-4703-8544-6a7864531b18.png)

    * Xem phần Replication có role:master, connected_slaves:2 là cài đặt thành công:

    ![image](https://user-images.githubusercontent.com/36092539/177723004-20750c38-31e4-4d85-915c-df41737622ba.png)

    * Kiểm tra service đã bật bằng lệnh : netstat -tnlp 

    ![image](https://user-images.githubusercontent.com/36092539/177723534-12495f75-e85d-44ac-b2a1-c7c03a86808b.png)

4. Kết nối golang với redis sentinel:
    * Truy cập từng container lấy từng ip bằng lệnh: ifconfig
    * Trong ví dụ : 

    redis_master1 : 172.31.0.2

    redis_slave   : 172.31.0.4

    redis_sentinel : 172.31.0.5
    * Truy cập vào container core_go: ``` docker exec -it  core_go bash ```

    * Sửa lại connect redis : 
     ``` 
        cd /home/app/redisSentinel
        vim connectRedisSentinel.go
     ```
    * Sử lại config như hình :

    ![image](https://user-images.githubusercontent.com/36092539/177724493-5404a552-4b16-4b0e-b437-c7fb5b734c16.png)
    
    * Chạy file để kiểm tra :  ``` go run connectRedisSentinel.go ```
    * Kết quả :

    ![image](https://user-images.githubusercontent.com/36092539/177724921-eecccd74-22b4-497c-b0df-9e9496200bbc.png)
    * Truy cập từng node redis kiểm tra kết quả được kết quả như hình :

    Node master :

    ![image](https://user-images.githubusercontent.com/36092539/177725227-4c778ddf-7e66-47cc-b10c-69f9866f4a15.png)

    2 node slave :

    ![image](https://user-images.githubusercontent.com/36092539/177725398-8691b01f-9488-4522-84a6-26c132d21274.png)





    