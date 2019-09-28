# Abstract factory pattern

## Participants

* AbstractFactory
    - declares an interface for operations that create abstract product objects.
* ConcreteFactory
    - implements the operations to create concrete product objects. 
* AbstractProduct
    - declares an interface for a type of product object. 
* ConcreteProduct
    - defines a product object to be created by the corresponding concrete factory.
    - implements the AbstractProduct interface. 
* Client
    - uses only interfaces declared by AbstractFactory and AbstractProduct classes.
