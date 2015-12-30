README
------

This is a simple micro service which returns the header information of the 
request it hails from. 

DONE
====
* Create the base case JSON marshaller and unmarshaller. 

TODO
====

* Add a request server frontend to it, listen for a handler 
  and pass the information on to the next service layer.
* Run apache benchmark on a single server instane.
* Scaling - Write behind a scalable backend.
