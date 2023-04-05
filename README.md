# quicq
message broker over quic



# Paper

## *Abstract*
    Очереди сообщений стали неотъемлемой частью построения сложных приложений на базе микросервисной архитектуры.
    Такие их представители как Kafka, RabbitMQ способны обрабатывать от нескольких тысяч до миллиона сообщений в секунду. 
    В данной статье представлена архитектура самописной очереди сообщений quicq, работающей поверх протокола quic, пришедшего 
    на замену tcp+tls. 
    Надо придумать аннотацию, но будет что-то похожее на смесь message queue, quic и моей реализации mq


## *Architecture*
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

