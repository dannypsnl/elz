Quick Start
===========

Variable
--------

We have two kinds of variable, global & local.

Global variable be defined at module level.

Here has a variable named **a**

.. code::

    a = 10

Then local? Local variable available in function, method, macro.

.. code::

    fn main() {
      let a = 10
    }

Function
--------

Function in Elz is a procedure start with keyword `fn`.

.. code::

    fn add(l, r: i32) -> i32 { return l + r }

This example show that Elz will trying to complete type if you do not define it. At here, `l` be completed by type of `r`.

Array
-----

.. code::

    fn main() {
      let a = [1, 2, 3]
      let b = [1; 10]
      let c = a[0]
    }

Code shows array has two ways to define, `[]` with expression list or `[expression; length]`.

And you can extract value from array. Notice that array in Elz is immutable, so you won't find push method for array.

Type of an array is `[T; len]`. `len` means length.

Type
----

Define your type in Elz using keyword `type` following type's name and a tuple.

.. code::

    type Person (
      age: i32,
    )

    fn main() {
      let danny = Person(20)
    }

Notice that in now version you have to fill whole type constructor function.
