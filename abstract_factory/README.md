# Abstract factory pattern

## UML

![abstract factory](../media/abstract_factory.jpg)

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

## Acknowlege

* UML image reference from [Wikipedia abstract factory pattern](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
* Participants reference from [Design patterns GoF]
