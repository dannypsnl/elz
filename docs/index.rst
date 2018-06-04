.. Elz documentation master file, created by
   sphinx-quickstart on Mon Jun  4 17:36:18 2018.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

Welcome to Elz's documentation!
===============================

.. toctree::
   :maxdepth: 2
   :caption: Contents:

Introduciton
============

Elz is a programming language for more productive.

Quick Start
===========

Variable
--------

We have two kinds of variable, global & local.

Global variable be defined at module level.

Here has a variable named **a**

.. code:: rust
    a = 10

Then local? Local variable available in function, method, macro.

.. code:: rust
    fn main() {
      let a = 10
    }

Function
--------

Function in Elz is a procedure start with keyword `fn`.

.. code:: rust
    fn add(l, r: i32) -> i32 { return l + r }

This example show that Elz will trying to complete type if you do not define it. At here, `l` be completed by type of `r`.

Array
-----

.. code:: rust
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

.. code:: rust
    type Person (
      age: i32,
    )

    fn main() {
      let danny = Person(20)
    }

Notice that in now version you have to fill whole type constructor function.

Indices and tables
==================

* :ref:`genindex`
* :ref:`modindex`
* :ref:`search`
