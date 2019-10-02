# Composite pattern

## UML

![composite](../media/composite.jpg)

## Defnition

Composite design pattern is target to treat hierarchical objects uniformly. Without using the composite pattern, the client need to treat the `leaf` node in one way, and treat the `container` of the leaves in another way. With composite pattern, the client just need to reference to one uniform abstract class.

* Component
  * Abstract `leaf` and `container` of the leaves in one abstraction class.
  * Client does not make references to the `leaf` and `container` now. So the logic of the client is very simple.
* Leaf
  * The leaf of the object tree. Subclass `Component` interface, and implement the interface
* Composite
  * The non-leaf node of the object tree. Will contain one or more leaves. When implementing the `Component` interface, need to call the `Operation` methods of the children recursively.

## Acknowlege

* UML image reference from [Wikipedia composite pattern](https://en.wikipedia.org/wiki/Composite_pattern)
