* Pub-Sub Model : 
    Publish/subscribe messaging, or pub/sub messaging, is a form of asynchronous service-to-service communication used in serverless and microservices architectures. In a pub/sub model, any message published to a topic is immediately received by all of the subscribers to the topic. Pub/sub messaging can be used to enable event-driven architectures, or to decouple applications in order to increase performance, reliability and scalability.
    Pub/Sub enables you to create systems of event producers and consumers, called publishers and subscribers. Publishers communicate with subscribers asynchronously by broadcasting events, rather than by synchronous remote procedure calls (RPCs).

    It is based on the Observer Design Pattern

* MSA:

    Microservices architecture (often shortened to microservices) refers to an architectural style for developing applications. Microservices allow a large application to be separated into smaller independent parts, with each part having its own realm of responsibility. To serve a single user request, a microservices-based application can call on many internal microservices to compose its response.

    #### A microservices architecture is a type of application architecture where the application is developed as a collection of services. It provides the framework to develop, deploy, and maintain microservices architecture diagrams and services independently.

    * 5 core components of microservices architecture: 
    1. Microservices:
        Microservices make up the foundation of a microservices architecture. The term illustrates the method of breaking down an application into generally small, self-contained services, written in any language, that communicate over lightweight protocols.
    
    2. Containers
        Containers are units of software that package services and their dependencies, maintaining a consistent unit through development, test and production. Containers are not necessary for microservices deployment, nor are microservices needed to use containers. However, containers can potentially improve deployment time and app efficiency in a microservices architecture more so than other deployment techniques, such as VMs.

        The major difference between containers and VMs is that containers can share an OS and middleware components, whereas each VM includes an entire OS for its use. By eliminating the need for each VM to provide an individual OS for each small service, organizations can run a larger collection of microservices on a single server.
    
    3. Service mesh
        In a microservices architecture, the service mesh creates a dynamic messaging layer to facilitate communication. It abstracts the communication layer, which means developers don't have to code in inter-process communication when they create the application.


    4. Service discovery
        Whether it's due to changing workloads, updates or failure mitigation, the number of microservice instances active in a deployment fluctuate. It can be difficult to keep track of large numbers of services that reside in distributed network locations throughout the application architecture.

        Service discovery helps service instances adapt in a changing deployment, and distribute load between the microservices accordingly.

        The service discovery component is made up of three parts:

        1. A service provider that originates service instances over a network;
        2. A service registry, which acts as a database that stores the location of available service instances; and
        3. A service consumer, which retrieves the location of a service instance from the registry, and then communicates with that instance.

    5. API gateway
        Another important component of a microservices architecture is an API gateway. API gateways are vital for communication in a distributed architecture, as they can create the main layer of abstraction between microservices and the outside clients. The API gateway will handle a large amount of the communication and administrative roles that typically occur within a monolithic application, allowing the microservices to remain lightweight. They can also authenticate, cache and manage requests, as well as monitor messaging and perform load balancing as necessary.

* Error vs Exception

    An Error "indicates serious problems that a reasonable application should not try to catch."
    An Exception "indicates conditions that a reasonable application might want to catch."

    Example of Error
>   public class ErrorExample {
        public static void main(String[] args){
            recursiveMethod(10)
        }
        public static void recursiveMethod(int i){
            while(i!=0){
                i=i+1;
                recursiveMethod(i);
            }
        }
    }


    Example of Exception
>   public class ExceptionExample {
        public static void main(String[] args){
            int x = 100;
            int y = 0;
            int z = x / y;
        }
    }