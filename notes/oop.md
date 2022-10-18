# The Array
- Stores many pieces of data

*CANNOT* store different types of data

# The Structure
- Stores many pieces of data 
- CAN store different types of data

*The main issue with structs is that you cannot define functions within one*

# Objects/Classes
- Objects are instances of a class
- Classes are templates for objects

# 4 main princples of OOP
- Encapsulation
- Abstraction
- Inheritance
- Polymorphism

## Encapsulation
- Encapsulation refers to building data with methods that can operate on that data within a class

Essentially, it is the idea of hiding data within a class, preventing anything outside that class from directly interacting with it.

Members of other classes can interact with the attributes of another object through its methods.

These methods are known as: 
- getting methods
Retrieving information 
- setting" methods
Changing information. Setting methods also allow the programmer to easily keep track of attributes that depend on one another
*Remember, methods are the functions defined within the class*

- Its generally best to not allow external classes to directly edit an object's attributes (Very important when working on larger and complex programs)
- Each piece should not have access to or rely on the inner workings of other sections of code

Encapsulation:
    - Keeps the programmer in control of access to data
    - Prevents the program from ending up in any strange or unwanted states

## Abstraction
- Abstraction refers to only showing essential details and keeping everything else hidden
- Abstraction allows the program to be worked on incrementally and prevents it from becoming entangled and complex
- Determine specific points of contact that can act as an interface between classes, and only worry about the implementation when coding it

The classes you create should act like your car. Users of your classes should not worry about the inner details of those classes
Classes should not directly interact with other classes' data.

Modern programs are very complex to the point where multiple programmers tend to work on one program
In this case, it's best if the section that you work on is able to function without knowledge of the inner workings of your colleagues section. 
To achieve this it's best to think of your projects as:
- Interface
    and
- Implementation

    ### Interface
    The interface refers to the way sections of code can communicate with one another.
    This typically is done through methods that each class is able to access.

    ### Implementation
    The implementation of these methods, or how these methods are coded, should be hidden.

If classes are entangled, then one change creates a ripple effect that causes many more changes. Creating an interface through which classes can interact ensures that each piece can be individually developed.


# Inheritance
- Inheritance is the principle that allows classes to derive from other classes

    ### Class hierarchy
    The class hierarchy acts as a web of classes with different relationships to one another.
    In most cases the class hierarchy you create will have many more layers with many more classes in each layer.
    Ex:
    + class Item
            |
        class Weapon, class Tool
            |
          class Sword, class Club
            |
            class Sword, class Sword

    ### Access Modifiers
    Access modifiers change which classes have access to other classes, methods, or attributes.
    Here are three common access modifiers are:
    - Public 
    - Private
    - Protected 

    Public members can be accessed from anywhere in your program. This includes anywhere both inside of the class hierarchy it is defined as well as outside in the rest of the program.

    Private members can only be accessed from within the same class that the member is defined. This allows you to create other private members of the same name in different locations so that they do not conflict with one another.

    Protected members can be accessed within the class it is defined, as well as any subclasses of that class.
    This essentially makes protected members private to the hierarchy in which they are defined.


# Polymorphism
- Polymorphism describes methods that are able to take on many forms

There are two types of polymorphism:
- Dynamic polymorphism
- Static polymorphism

    ### Dynamic polymorphism
    Occurs during the runtime of the program. This type of polymorphism descibes when a method signature is in both a subclass and a superclass.

    The methods share the same name but have different implementation. 
    The implementation of the subclass that the object is an instance of overrides that of the superclass.

    ### Static polymorphism
    Static polymorphism occurs during compile-time rather than during runtime. This refers to when multiple methods with the same name but different arguments are defined in the same class.

    Ways to differentiate methods of the same name:
    - Different number of parameters
    - Different types of parameters
    - Different order of parameters

    This is known as method overloading. Despite the methods having the same name, their signatures are different due to their different arguments.

    Keep in mind that method overloading can cause trouble  if you do not keep straight which parameters you need for which implementation.
