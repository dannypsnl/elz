# Change log

### 0.1.0: in progress

#### Syntax

* global variable
    ```elz
    x: int = 1;
    ```
* global function definition
    ```elz
    main(): void {}
    foo(): void {}
    ```
* local variable
    ```elz
    main(): void {
      x: int = 1;
    }
    ```
* function call
    ```elz
    main(): void {
      println("hello, world");
    }
    ```
* string template
    ```elz
    main(): void {
      x: int = 1;
      println("x = {x}");
    }
    ```
* List literal
    ```elz
    x: List[int] = [];
    ```

#### Type

* `void`
* `int`
* `string`
* `bool`
* `f64`
* `List[T]`
* function type, e.g. `(int, int): int`