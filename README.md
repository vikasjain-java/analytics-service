================================================ Problem ================================================
Golang program should be able to track the unique visitors visiting a given URL.

The service should have two endpoints: 
1. One to capture visitor information.
2. One to query the number of unique visitors per URL

For example, if there are two users, Alice and Bob:
- Alice visited: 
  - http://foo.com
  - http://foo.com
  - http://bar.com

- Bob visited:
  - http://bar.com?query=q

The query endpoint should provide the following information:
  - http://foo.com, 1 visitor
  - http://bar.com, 2 visitors

================================================ Solution ================================================ 

2 endpoints are exposed and server run on port 8080

server url - http://localhost:8080

1) Endpoint which capture the visit of person can be hit with below url and param (I tried using postman)
    http://localhost:8080/capture?url=http://bar.com&visitorName=bob
2) Endpoint which used to getAllVistors hit different url
    http://localhost:8080/query
3) Endpoint which gives hit to particular url 
    http://localhost:8080/query?url=http://bar.com

Implementation has two function for query and capture.
It has one struct which contains -
    map of map (this stores url and userName)
    mutex to get lock 
