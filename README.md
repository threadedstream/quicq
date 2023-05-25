# quicq
message broker over quic



# Paper

Есть идея назвать paper Quic + Message Queue = quicq или Очередь сообщений поверх протокола quic  

## Abstract (Аннотация)
    

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
    
    **Congestion control**
    Another problem is that we don’t know up front how much the maximum bandwidth will be. It often depends on a bottleneck somewhere in the end-to-end connection, but we cannot predict or know where this will be. The Internet also doesn’t have mechanisms (yet) to signal link capacities back to the endpoints.

    Additionally, even if we knew the available physical bandwidth, that wouldn’t mean we could use all of it ourselves. Several users are typically active on a network concurrently, each of whom need a fair share of the available bandwidth.

    As such, a connection doesn’t know how much bandwidth it can safely or fairly use up front, and this bandwidth can change as users join, leave, and use the network. To solve this problem, TCP will constantly try to discover the available bandwidth over time by using a mechanism called congestion control.

    At the start of the connection, it sends just a few packets (in practice, ranging between 10 and 100 packets, or about 14 and 140 KB of data) and waits one round trip until the receiver sends back acknowledgements of these packets. If they are all acknowledged, this means the network can handle that send rate, and we can try to repeat the process but with more data (in practice, the send rate usually doubles with every iteration).

    **Very important**
    They observed QUIC outperforms TCP/TLS with HTTP/2 for wireless and mobile networks, whereas for wired and stable networks, a significant performance improvement is not seen [29].

    Write that QUIC is also meant to avoid MITM attacks, as it encrypts both headers and payload

## III. Message Queues (Очереди сообщений) 

## General architecture (Обобщенная архитектура)
    Написать про очереди сообщений, вариации, архитектуру и области применения. 

### IV. Архитектура quicq 

## quicq Architecture
Условно quicq можно разделить на клиентскую и серверную части. Клиентская часть представляет собой сущность, которая пишет сообщения (далее продюсер), а также сущность, которая эти сообщения читает (далее консьюмер). Серверная часть представлена брокером - сущностью, выполняющую всю работу по распределению сообщений в специальные хранилища (топик), а также передачу этих сообщений клиентам. 
Топик представляет собой очередь, хранящую сообщения в FIFO манере. Очереди могут существовать как в непартицированном, так и в партицированном виде. Очередь в партицированном виде позволяет снизить нагрузку на саму очередь в целом, распределяя сообщения по очередям по определенному критерию. 
Необходимо позволять конечным пользователям создавать свои реализации алгоритмов партицирования. 

Подключение к брокеру
Для подключения к брокеру клиент отправляет ClientConnect запрос с указанием следующих опций: 
1. ClientID - айди продюсера или консьюмера
2. MessageDeliveryPolicy - (
    0 - at least once
    1 - at most once
    2 - exactly once 
)

При успешном подключении сервер отправляет в ответ ClientConnectAck с ID подключившегося клиента. 

Протокол
Общение с брокером предполагает наличие некоторой договоренности, называемой протоколом. Предполагается, что quicq будет поддерживать сразу несколько протоколов для эффективной передачи данных, но в этой секции затрагивается только один, являющийся нативным для quicq. 
Для отправки сообщения продюсер отправляет единственный Post запрос, в котором указывается имя топика, в который сообщение должно быть записано, а также само сообщение в виде строки. Брокер же должен создать топик в случае, если последний не был найден в списке существующих. В случае успешной записи сообщения в топик брокер должен послать ответ с именем топика, а также с идентификатором созданного сообщения. Для получения информации о топиках необходимо отправить запрос FetchTopicMetadata, ответ на который содержит такую информацию, как имя топика, размер очереди, а также дата создания топика. Для удаления топика необходимо отправить DestroyTopic запрос с указанием имени топика. 
Для чтения сообщений консьюмер подписывается на топик, предварительно отправив Subscribe запрос с указанием имени топика. Для отписки от топика отправляется, соответственно, Unsubscribe запрос. В случае успешной подписки или отписки от топика брокер возвращает его метаданные. При чтении сообщений консьюмер отправляет Poll запрос с указанием ID консьюмера, отправляющего запрос. В ответ брокер отправляет список релевантных записей. Запись в quicq содержит идентификатор записи, ключ, полезную нагрузку и время создания. В сценариях выше описано лишь поведение брокера при благоприятных обстоятельствах, но нужно иметь в виду, что могут возикать проблемы во время исполнения этих запросов. В случае ошибки брокер должен вернуть объект ошибки, содержащий в себе тип ошибки и ее детали в виде строки. Каждый запрос содержит преамбулу и тело запроса. Преамбула включает в себя версию протокола и тип запроса. 

Режим подтверждения отправленных сообщений
quicq поддерживает 3 типа подтверждения сообщений: at least once, at most once, exactly once. В первом случае брокер уведомляет продюсера о статусе отправленного сообщения, во втором же случае продюсер отправляет сообщение, не ожидая никакого подтверждения. Последнее может быть полезно, когда важна скорость, а состояние отправленным сообщений не является приоритетом. 

Экспериментальные опции
Одной из эспериментальных опций является мультипротокольная коммуникация между сервером и клиентом. Из-за сложности реализации и возможного наличия большого количества проблем данная опция не является обязательной к реализации. На ранней стадии возможна реализация с использованием декодировщиков со стороны сервера и клиента, при которой принимающая сторона считывает название протокола и его версию в запросе и переводит это сообщение в язык протокола, с которым может работать принимающая сторона. 

ClientConnect 
ClientConnectAck
Post
PostAck
Poll 
PollResponse
FetchTopicMetadata
TopicMetadataResponse
Subscribe 
SubscribeAck
Unsubscribe
UnsubscribeAck 


### Conclusion (Заключение)

head of line blocking 
rabbitmq has an idea of channels which resembles http2/streams 

### Experiments
#### 1P1C 
Broker - 3% CPU, 0.1% MEM on avg
Consumer - 1% CPU, 0.1% MEM on avg
Producer - 1% CPU, 0.1% MEM on avg


### TODOs

- [x] Try, for the sake of experiment, sending huge number of data through tcp and QUIC pipes. Evaluate performance   
- [x] Find comparison graphs of quic cubic and tcp cubic \
- [x] spdy - read about that as well
- [x] Add a global config with parameters like broker address, queue length and etc.. \
- [ ] Define ErrorResponse object to report errors to clients \
- [ ] Add more brokers to the scene and implement leader election mechanism? \
- [ ] Properly close connection  

write down your pc specs you ran tests on

Subscribe()
Unsubscribe()
Post()
Poll() 

Subscribe()

subscriber -> Topic

Попробовать мультипротокольную реализацию с поддержкой нескольких протоколов, при которой клиент и сервер используют разные протоколы, но понимают друг друга

**Bibliography**

Prashant Kharat, Muralidhar Kulkarni, Modified QUIC protocol with congestion control for improved network performance, IET Communications https://ietresearch.onlinelibrary.wiley.com/doi/10.1049/cmu2.12154, Vol.5, Issue 9 1210-1222.

Robin Marx, 2021, https://www.smashingmagazine.com/2021/08/http3-performance-improvements-part2.

Puneet Kumar, QUIC - A Quick Study, arxiv.org, https://arxiv.org/pdf/2010.03059.pdf

Puneet Kumar and Behnam Dezfouli. Implementation and analysis of quic for mqtt.
Computer Networks, 150:28–45, 2019, https://arxiv.org/pdf/1810.07730.pdf

Репозиторий с кодом сервера RabbitMQ, https://github.com/rabbitmq/rabbitmq-server

AMQP по-русски, 2009, https://habr.com/ru/articles/62502/ 

Philippe Dobbelaere, Kyumars Sheykh Esmaili, Kafka versus RabbitMQ, 2017, arxiv.org, https://arxiv.org/pdf/1709.00333.pdf

https://docs.confluent.io/kafka/design/delivery-semantics.html

https://dev.to/behalf/event-ordering-with-apache-kafka-2gb5

RabbitMQ, Reliability Guide, https://www.rabbitmq.com/reliability.html

RabbitMQ, Broker Semantics, https://www.rabbitmq.com/semantics.html

Google,Protocol Buffers, https://protobuf.dev - 

Репозиторий с кодом quicq, https://github.com/threadedstream/quicq 


