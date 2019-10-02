# Facade pattern

## UML

![facade](../media/facade.jpg)

## Defnition

In a complicated system, it will always contains serval sub systems. In point of view of the client, it will need to know every detail of the sub systems to use the services provided by them. So the client need to tightly couple with the sub systems. Within `facade pattern`, `facade` class will hide the detail of the subsystems, and expose more simple interfaces to client.

* Subsystem
  * Business function related class
  * May contain lot of interfaces, and not so easy to use
* Facade
  * The high level wrapper class of the related subsystems, and abstract some simple interfaces
  * Easy to use
  * decoupling the subsystems from the client

## Acknowlege

* UML image reference from [Wikipedia facade pattern](https://en.wikipedia.org/wiki/Facade_pattern)
