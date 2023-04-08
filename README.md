# quicq
message broker over quic



# Paper

Есть идея назвать paper Quic + Message Queue = quicq или Очередь сообщений поверх протокола quic  

## *Abstract*
    Очереди сообщений стали неотъемлемой частью построения сложных приложений на базе микросервисной архитектуры.
    Такие их представители как Kafka, RabbitMQ способны обрабатывать от нескольких тысяч до миллиона сообщений в секунду. 
    В данной статье представлена архитектура самописной очереди сообщений quicq, работающей поверх протокола quic, пришедшего 
    на замену tcp+tls. 
    Надо придумать аннотацию, но будет что-то похожее на смесь message queue, quic и моей реализации mq

## *Quic*
    Написать про замену tcp под названием quic. Перечислить главные недостатки tcp и указать на достоинства quic 

## *Message Queues*
    Написать про очереди сообщений, вариации, архитектуру и области применения. 


## *quicq Architecture*
    The idea of consumers and producers reading from/writing to queue, queue must be thread-safe to avoid panics 
    We still have the idea of committed entries
    We also need a broker to nicely manage communication between producers and consumers 
    Need to think over a proper way to implement a queue.
    The retention policy to know when to purge all messages

    Use kafka records as entries?

    *What are the responsibilities?*
    Consumer - reads messages
    Producer - produces messages
    Broker - handles all communication stuff, like connecting to a server, etc...

    Здесь надо будет дать представление о том, что quicq из себя представляет. 

head of line blocking 
rabbitmq has an idea of channels which resembles http2/streams 

spdy - read about that as well

rabbitmq 
uses channels
uses tcp streams
idea of publishers and consumers
multiplexing  

maybe turn to pub/sub?
implement a custom pub/sub protocol over quic 

requests/response

Subscribe()
Unsubscribe()
Post()
Poll() 

https://github.com/rabbitmq/rabbitmq-server
https://habr.com/ru/articles/62502/