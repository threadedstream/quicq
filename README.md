# quicq
message broker over quic



# Paper

Есть идея назвать paper Quic + Message Queue = quicq или Очередь сообщений поверх протокола quic  

## Abstract (Аннотация)
    Очереди сообщений стали неотъемлемой частью построения сложных приложений на базе микросервисной архитектуры.
    Такие их представители как Kafka, RabbitMQ способны обрабатывать от нескольких тысяч до миллиона сообщений в секунду. 
    В данной статье представлена архитектура самописной очереди сообщений quicq, работающей поверх протокола quic, пришедшего 
    на замену tcp+tls. 
    Надо придумать аннотацию, но будет что-то похожее на смесь message queue, quic и моей реализации mq

## I. Introduction (Введение)
    

    Статья организована следующим образом: во второй главе ведется описание архитектуры протокола quic, объясняется причина его возникновения, а также отмечается ряд достоинств, делающий quic более предпочтительным по сравнению с tcp, во третьей главе 
    приводится обобщенное описание очередей сообщений, их различие, а также приводится ряд примеров реализаций, используемых в приложениях с миллионами пользователей, в конце статьи приводится заключение проделанной работы.  

## II. Networking preliminaries (Сеть)

### Quic
    Разработанный в Google протокол quic (англ. Quick UDP Internet Connections) пришел на замену схеме tcp+tls, имеющей ряд недостатков влияющих на скорость как регистрации соединения, так и передачи данных. 
    Одним из недостатков tcp является его механизм контроля перегрузки (англ. Congestion Control), заключающийся в 
    постепенном увеличении объема передаваемых данных. В качестве начального значения объема передаваемых данных чаще всего выступает значение в 14 килобайт. Существует большое количество алгоритмов, имеющих в своих реализациях оптимизации, позволяющие ускорить передачу данных, такие как Cubic, BIC, Westwood, NewReno и т.д. В реализации QUIC также предусмотрен механизм контроля перегрузки - преимуществом можно назвать относительную простоту экспериментирования и более широкую аудиторию 
    исследователей по той причине, что tcp реализован на уровне ядра операционной системы, тогда как QUIC реализован на пользовательском уровне.  

    ** QUIC is still bound by the laws of physics and the need to be nice to other senders on the Internet. This means that it will not magically download your website resources much more quickly than TCP. However, QUIC’s flexibility means that experimenting with new congestion-control algorithms will become easier, which should improve things in the future for both TCP and QUIC.**
    
    **QUIC is still bound by the laws of physics and the need to be nice to other senders on the Internet. This means that it will not magically download your website resources much more quickly than TCP. However, QUIC’s flexibility means that experimenting with new congestion-control algorithms will become easier, which should improve things in the future for both TCP and QUIC.**

    **Congestion control**
    Another problem is that we don’t know up front how much the maximum bandwidth will be. It often depends on a bottleneck somewhere in the end-to-end connection, but we cannot predict or know where this will be. The Internet also doesn’t have mechanisms (yet) to signal link capacities back to the endpoints.

    Additionally, even if we knew the available physical bandwidth, that wouldn’t mean we could use all of it ourselves. Several users are typically active on a network concurrently, each of whom need a fair share of the available bandwidth.

    As such, a connection doesn’t know how much bandwidth it can safely or fairly use up front, and this bandwidth can change as users join, leave, and use the network. To solve this problem, TCP will constantly try to discover the available bandwidth over time by using a mechanism called congestion control.

    At the start of the connection, it sends just a few packets (in practice, ranging between 10 and 100 packets, or about 14 and 140 KB of data) and waits one round trip until the receiver sends back acknowledgements of these packets. If they are all acknowledged, this means the network can handle that send rate, and we can try to repeat the process but with more data (in practice, the send rate usually doubles with every iteration).

    **TODO(me)**: Написать про замену tcp под названием quic. Перечислить главные недостатки tcp и указать на достоинства quic 

    **Very important**
    They observed QUIC outperforms TCP/TLS with HTTP/2 for wireless and mobile networks, whereas for wired and stable networks, a significant performance improvement is not seen [29].

    Write that QUIC is also meant to avoid MITM attacks, as it encrypts both headers and payload

## III. Message Queues (Очереди сообщений) 

## General architecture (Обобщенная архитектура)
    Написать про очереди сообщений, вариации, архитектуру и области применения. 

### IV. Архитектура quicq 

## quicq Architecture
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

### Conclusion (Заключение)

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

**Bibliography**

Prashant Kharat, Muralidhar Kulkarni, Modified QUIC protocol with congestion control for improved network performance, IET Communications https://ietresearch.onlinelibrary.wiley.com/doi/10.1049/cmu2.12154, Vol.5, Issue 9 1210-1222.

Robin Marx, 2021, https://www.smashingmagazine.com/2021/08/http3-performance-improvements-part2.

Puneet Kumar, QUIC - A Quick Study, arxiv.org, https://arxiv.org/pdf/2010.03059.pdf

Репозиторий с кодом rabbitmq сервера, https://github.com/rabbitmq/rabbitmq-server

AMQP по-русски, 2009, https://habr.com/ru/articles/62502/ 
