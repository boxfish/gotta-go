# Interfaces
1. Interface types express generalizations or abstractions about the behaviors of other types. It lets us write functions that are more flexible and adaptable because they are not tied to the details of one particular implementation.
2. what makes Go’s interfaces so distinctive is that they are satisfied implicitly. In other words, there’s no need to declare all the interfaces that a given concrete type satisfies; simply possessing the necessary methods is enough.

## Interfaces as Contract
1. This freedom to substitute one type for another that satisfies the same interface is called substitutability, and is a hallmark of object-oriented programming.

## Interface Types
1. We can declare new interface types as combinations of existing ones.

## Interface Satisfaction
1. a value of type T does not possess all the methods that a *T pointer does, and as a result it might satisfy fewer interfaces.
2. Only the methods revealed by the interface type may be called, even if the concrete type has others:

## Interface Values
1. A value of an interface type, or interface value, has two components, a concrete type and a value of that type. These are called the interface’s dynamic type and dynamic value.
2. In general, we cannot know at compile time what the dynamic type of an interface value will be, so a call through an interface must use dynamic dispatch. Instead of a direct call, the compiler must generate code to obtain the address of the method named Write from the type descriptor, then make an indirect call to that address. The receiver argument for the call is a copy of the interface’s dynamic value.
3. Interface values may be compared using == and !=. Two interface values are equal if both are nil, or if their dynamic types are identical and their dynamic values are equal according to the usual behavior of == for that type. Because interface values are comparable, they may be used as the keys of a map or as the operand of a switch statement.
4. However, if two interface values are compared and have the same dynamic type, but that type is not comparable (a slice, for instance), then the comparison fails with a panic.
5. A nil interface value, which contains no value at all, is not the same as an interface value containing a pointer that happens to be nil.

## Sorting with sort.Interface
1. An in-place sort algorithm needs three things—the length of the sequence, a means of com- paring two elements, and a way to swap two elements—so they are the three methods of sort.Interface: `Len() int`, `Less(i, j int) bool`, `Swap(i, j int)`.

## The http.Handler Interface
1. If we need to wrap a function into an interfce with sole method that has the same funciton signature, we can define it as a function type and the behavior of the method is just to call the undelrying function.

## The error interface
1. In `errors` package, the `New` funciton returns a pointer to the underlying error struct so that each error created won't compare equal to others.
