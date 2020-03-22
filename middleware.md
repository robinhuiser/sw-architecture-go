# Middleware

The book outlines the following middleware to be installed via `homebrew`:

~~~bash
# Install formulas
$ brew install mysql cassandra redis kafka
~~~

What... no NSQ? If you do a `brew install nsq`, and then a `brew services start nsq`, you end up in a weird place NSQ-wise...

While `nsqd` will be running, you won't have an `nsqlookupd`. So most of the NSQ examples you find online won't work. If you start an nsqlookupd separately, it still won't work, because the plist file created by the default formula doesn't start nsqd with the necessary arguments to find and use `nsqlookupd`, hence:

~~~bash
$ brew install btubbs/nsq/nsqlookup
~~~

## Starting / stopping

Any of these middleware can be started or stopped as a **service** using: `brew services start [formula]`:

~~~bash
# Start (copy-paste snippet)
brew services start mysql
brew services start cassandra
brew services start redis
brew services start zookeeper
brew services start kafka
brew services start btubbs/nsq/nsqlookup
brew services start btubbs/nsq/nsq

# Stop (copy-paste snippet)
brew services stop mysql
brew services stop cassandra
brew services stop redis
brew services stop kafka
brew services stop zookeeper
brew services stop btubbs/nsq/nsq
brew services stop btubbs/nsq/nsqlookup
~~~

 or as **standalone** via the commands below:

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

# NSQ (2 processes)
$ nsqlookupd
$ nsqd --data-path=/usr/local/var/nsq --lookupd-tcp-address=127.0.0.1:4160
~~~

## GUIs

I am using the following tools to access the middleware services:

* [TablePlus](https://tableplus.com/) to access Redis, MySQL and Cassandra
* [Kafka Tool](http://www.kafkatool.com/) for Kafka/Zookeeper
* [NSQAdmin](http://localhost:4171) via `nsqadmin --lookupd-http-address=127.0.0.1:4161`
