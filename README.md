# Software architecture with Golang

My notes, thoughts and code working through "Hands-on software architecture with Golang"

## Prerequisites

The book outlines middleware to be installed via `homebrew`:

~~~bash
# Install formulas
$ brew install mysql cassandra redis kafka nsq
~~~

Note: Any of these middleware can be started as a **service** using: `brew services start [formula]` or as **standalone** via the commands below:

~~~bash
#
# Mysql server
$ mysql_secure_installation   # Secure the instance
$ mysql -uroot                # Allows connections from localhost by default
$ mysql.server start

#
# Cassandra
$ cassandra -f

#
# Redis
$ redis-server /usr/local/etc/redis.conf

#
# Zookeeper & Kafka
$ zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties & \
   kafka-server-start /usr/local/etc/kafka/server.properties

#
# NSQ
$ nsqd -data-path=/usr/local/var/nsq        # Service
$ nsqadmin -http-address 127.0.0.1:4171     # Web interface on http://localhost:4171
~~~
