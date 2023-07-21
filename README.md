# flight

This project receives a list of flight routes, orders and returns a list of possible routes within the requested route.

Payload:
```json
[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]]
```

Output:
```json
{"meta":{"success":true,"error":[]},"data":[[["DUB","LHR"]],[["DUB","LHR"],["LHR","GVA"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["LHR","GVA"]],[["LHR","GVA"],["GVA","MXP"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["GVA","MXP"]],[["GVA","MXP"],["MXP","NCE"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["MXP","NCE"]],[["MXP","NCE"],["NCE","MAD"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["NCE","MAD"]],[["NCE","MAD"],["MAD","LIM"]],[["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["MAD","LIM"]],[["MAD","LIM"],["LIM","SCL"]],[["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["LIM","SCL"]],[["LIM","SCL"],["SCL","AEP"]],[["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["SCL","AEP"]],[["SCL","AEP"],["AEP","EZE"]],[["AEP","EZE"]]]}
```

### Examples

### cmd/benchmark

benchmark from de code

### cmd/localDevOps

The basis of this module is the chaos/failure framework. Made for testing microservices before the microservice goes in 
production server.

It was made to create all the necessary infrastructure for the microservice to work and then cause random failures, such 
as stopping the container or simulating an overloaded network.

This project will be create files in the local computer with security breach reports and resource consumption, such as 
memory and cpu

To learn more about him https://github.com/helmutkemper/chaos

Project structure:

```
	//
	// +-------------+     +-------------+     +-------------+
	// |             |     |             |     |             |
	// |  Docker  0  |     |  Docker  1  |     |  Docker  2  |
	// |   Server    |     |   Server    |     |   Server    |
	// |             |     |             |     |             |
	// +------+------+     +------+------+     +------+------+
	//        ↓                   ↓                   ↓
	// -------+---------+--Docker--Network--+---------+-------
	//                            ↑       
	//                     +------+------+
	//                     |             |
	//                     |   Reverse   |
	//                     |    Proxy    |
	//                     |             |
	//                     +------+------+
	//
```

This was the first time I did the reverse proxy chaos/failure test and I discover that it doesn't work in the container.
However, this is a project made when I was learning Golang, about 6 years ago. (https://github.com/helmutkemper/basicReverseProxy)

I know this is my fault, but somehow the chaos/failure test did its job in detecting the problem.

Note: after a few minutes, server containers will start to fail.

### cmd/server

This is the project server. It has an endpoint `http://localhost:8080/calculate`

Payload:
```json
[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]]
```

Output:
```json
{"meta":{"success":true,"error":[]},"data":[[["DUB","LHR"]],[["DUB","LHR"],["LHR","GVA"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["DUB","LHR"],["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["LHR","GVA"]],[["LHR","GVA"],["GVA","MXP"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["LHR","GVA"],["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["GVA","MXP"]],[["GVA","MXP"],["MXP","NCE"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["GVA","MXP"],["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["MXP","NCE"]],[["MXP","NCE"],["NCE","MAD"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["MXP","NCE"],["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["NCE","MAD"]],[["NCE","MAD"],["MAD","LIM"]],[["NCE","MAD"],["MAD","LIM"],["LIM","SCL"]],[["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["NCE","MAD"],["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["MAD","LIM"]],[["MAD","LIM"],["LIM","SCL"]],[["MAD","LIM"],["LIM","SCL"],["SCL","AEP"]],[["MAD","LIM"],["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["LIM","SCL"]],[["LIM","SCL"],["SCL","AEP"]],[["LIM","SCL"],["SCL","AEP"],["AEP","EZE"]],[["SCL","AEP"]],[["SCL","AEP"],["AEP","EZE"]],[["AEP","EZE"]]]}
```
