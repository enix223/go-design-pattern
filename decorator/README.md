# Decorator pattern

## UML

![decorator](../media/decorator.jpg)

## Defnition

When client need to extend the function of existing module, and we don't want to change the module or we can't change the module, we can achive our goal by `decorator pattern`.

* Component
  * The existing business logic related interface.
* Component1
  * The existing class which implement the `Component` interface
  * The class we need to extend
* Decorator1 & Decorator2
  * An new class which using composition to hold a reference to the `Component` interface, and adding some new logic to the `Operation` method
  * Decorator call the new code and the `component.Operation` to extend the `Component` interface.

## Acknowlege

* UML image reference from [Wikipedia decorator pattern](https://en.wikipedia.org/wiki/Decorator_pattern)
