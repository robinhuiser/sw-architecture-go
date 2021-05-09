# Software architecture with Golang

My notes, thoughts and code working through "Hands-on software architecture with Golang"

## Prerequisites

First, I [installed the required middleware](./install-middleware.md) on my MacBook Pro.

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

## Chapter notes

* [Chapter 1 - Building Big with Go](./01_building_big_with_go/README.md)
* [Chapter 2 - Packaging Code](./02_packaging_code/README.md)
* [Chapter 3 - Design Patterns](./03_design_patterns/README.md)
