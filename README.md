README
------

This is a simple micro service which returns the header information of the 
request it hails from. 

DONE
====
* Create the base case JSON marshaller and unmarshaller. 
* Add a request server frontend to it, listen for a handler 
  and pass the information on to the next service layer.

TODO
====

* Run apache benchmark on a single server instane.
* Use goroutines to speed it up. 
* Add a load balancer layer in front of it. 
* Scaling - Write behind a scalable backend.

RUN
===

$ go build
$ ./scalable-rest
