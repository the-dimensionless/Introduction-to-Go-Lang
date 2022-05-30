* What is meant by Ellipsis?

    In computer programming, ellipsis notation (.. or ...) is used to denote ranges, an unspecified number of arguments, or a parent directory. Most programming languages require the ellipsis to be written as a series of periods; a single (Unicode) ellipsis character cannot be used.
    ex -> function f(int a, ...args)


* what is meant by Variadic functions?

    In mathematics and in computer programming, a variadic function is a function of indefinite arity, i.e., one which accepts a variable number of arguments. Support for variadic functions differs widely among programming languages.
    Variadic function adds flexibility to the program. It takes one fixed argument and then any number of arguments can be passed. The variadic function consists of at least one fixed variable and then an ellipsis(…) as the last parameter.

* Concurrency vs Parallelism

    Concurrency - Concurrency is about to handle numerous tasks at once. This means that you are working to manage numerous tasks done at once in a given period of time. However, you will only be doing a single task at a time. This tends to happen in programs where one task is waiting and the program determines to drive another task in the idle time. It is an aspect of the problem domain — where your program needs to handle numerous simultaneous events.

    Parallelism - Parallelism is about doing lots of tasks at once. This means that even if we have two tasks, they are continuously working without any breaks in between them. It is an aspect of the solution domain — where you want to make your program faster by processing different portions of the problem in parallel.

    #### Concurrency is NOT Parallelism

    A concurrent program has multiple logical threads of control. These threads may or may not run in parallel. A parallel program potentially runs more quickly than a sequential program by executing different parts of the computation simultaneously (in parallel). It may or may not have more than one logical thread of control.
 
* Installing Go from Source?

* ORM [Object Relational Mapping] :: ORM Library

    Object-relational mapping (ORM) is a programming technique in which a metadata descriptor is used to connect object code to a relational database. Object code is written in object-oriented programming (OOP) languages such as Java (Hibernate/iBatiss) or C#. ORM converts data between type systems that are unable to coexist within relational databases and OOP languages.

    ORM resolves the object code and relational database mismatch with three approaches: bottom up, top-down and meet in the middle. Each approach has its share of benefits and drawbacks. When selecting the best software solution, developers must fully understand the environment and design requirements.

    In addition to the data access technique, ORM's benefits also include:

    * Simplified development because it automates object-to-table and table-to-object conversion, resulting in lower development and maintenance costs
    * Less code compared to embedded SQL and handwritten stored procedures
    * Transparent object caching in the application tier, improving system performance
    * An optimized solution making an application faster and easier to maintain


* REST & API's : Main concepts & End Points

    REST is short for Representational State Transfer, an architectural style for building web services that interact via an HTTP protocol. Its principles were formulated in 2000 by computer scientist Roy Fielding and gained popularity as a scalable and flexible alternative to older methods of machine-to-machine communication. It still remains the gold standard for public APIs.

    The key elements of the REST API paradigm are

    * client or software that runs on a user’s computer or smartphone and initiates communication;
    * server that offers an API as a means of access to its data or features; and
    * resource, which is any piece of content that the server can provide to the client (for example, a video or a text file).
    To get access to a resource, the client sends an HTTP request. In return, the server generates an HTTP response with encoded data on the resource. Both types of REST messages are self-descriptive, meaning they contain information on how to interpret and process them.

    #### Any REST request includes four essential parts: an HTTP method, an endpoint, headers, and a body.

    1. An HTTP method describes what is to be done with a resource. There are four basic methods also named CRUD operations: GET PUT POST DELETE

    2. An endpoint contains a Uniform Resource Identifier (URI) indicating where and how to find the resource on the Internet. The most common type of URI is a Unique Resource Location (URL), serving as a complete web address.

    3. Headers store information relevant to both the client and server. Mainly, headers provide authentication data — such as an API key, the name or IP address of the computer where the server is installed, and the information about the response format.

    4. A body is used to convey additional information to the server. For instance, it may be a piece of data you want to add or replace.