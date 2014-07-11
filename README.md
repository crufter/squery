squery
======

***WHAT YOU READ HERE IS PROBABLY WRONG. YOU HAVE BEEN WARNED. Please if you know an existing implementation tell me, because though I love programming I love sleeping more.***

What is squery?
---
- A microservices protocol and library which helps you query your services as easily as you query the dom with jquery, or your database with SQL*

*Almost correct, keep reading

The problems
---
A microservices based architecture has plenty of benefits, but unfortunately it has drawbacks as well: it is quite hard to do more complex queries involving multiple services. This project helps you keep your services lousely coupled.

The solution
---
Crawl the data, copy it and precompute your views! This way you can ask your system high level questions without sacrificing performance, violating service boundaries! 

In action
---

Once the services define their schemas for the entities, and provide a way to be crawled, we can use a library to query the data in the following (nothing is set in stone) fashion:

Querying customers of company with id 23
```go
api.Query("customer").RefersTo(api.Query("company").Id(23))
```

Querying the financial records of a customer for a given month
```go
api.Query("finance.entry").RefersTo(api.Query("customer").Id(97)).Between("timeStamp", monthBegin, monthEnd)
```

These queries can be analyzed, and a view can be precomputed to return the data efficiently exactly the way you want. This lets you express your requirement in a high level fashion.
