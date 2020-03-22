# Software architecture with Golang

My notes, thoughts and code working through "Hands-on software architecture with Golang"

## Prerequisites

First, I [installed the required middleware](./middleware.md) on my MacBook Pro.

Next, I added the code examples to my repo under `./packt.com` as a submodule:

~~~bash
$ git submodule add \
   https://github.com/PacktPublishing/Hands-On-Software-Architecture-with-Golang
~~~

Last, I downloaded the color images for the book under `./packt.com` via: 

~~~bash
$ curl -XGET \
   'https://static.packt-cdn.com/downloads/9781788622592_ColorImages.pdf' \
   -o ./packt.com/9781788622592_ColorImages.pdf
~~~
