# Facade pattern

## UML

![flyweight](../media/flyweight.jpg)

## Defnition

In a system with sophisticated class, we offen don't want to recreate such object twice, and its states are immutable, so it can be resue through the whole system.

* Flyweight
  * An interface defining the business operations
* ConcretFlyweight
  * A class implement the business operations, and it is immutable, which can be resue through the whole system
* FlyweightFactory
  * A factory which create the flyweight object when it is not exist, or return the existing one.

## Acknowlege

* UML image reference from [Wikipedia flyweight pattern](https://en.wikipedia.org/wiki/Flyweight_pattern)


## Ref

[What are the practical use differences between Flyweight vs Singleton patterns?](https://stackoverflow.com/questions/16750758/what-are-the-practical-use-differences-between-flyweight-vs-singleton-patterns)
