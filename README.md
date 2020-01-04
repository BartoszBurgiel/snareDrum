# SnareDrum
SnareDrum is a program helping users to understand the concept of a turing machine and theoretical computer science's principles. Essentialy it's an advanced <a href="https://en.wikipedia.org/wiki/Brainfuck">brainfuck</a> interpreter where the user is able to create his own "version" of brainfuck with own tokens. 

#### With SnareDrum you can: 
* **Run** a program in **your own programming language**
* **Compile** your program into a binary file that can be executed later
* **Debug** your code by printing all of the stack as a table 
* "**Translate**" any text into an executable program (binary, or source code in **your** language)

</br>

## How tos

### Console commands 
* `build <path-to-project>`: write out binary file 
* `debug <path-to-project>`: print the entire stack 
* `exec <path-to-project/binary>`: execute binary/sd project file 
* `generate <path-to-lang> <text>`: generate program in given lang printing following text


### SnareDrum project file structure

<pre><code>project-name
  - header.json  <- language meta data
  - main.sd      <- source code
</code></pre>

#### header.json example [pure brainfuck]

<pre><code>{
    "_comment": "Brainfuck lang setup",
    "SingleChard": true,
    "Pointer": {
        "Up": ">",
        "Down": "<"
    },
    "IO": {
        "In": ",",
        "Out": "."
    },
    "Cell": {
        "Add": "+",
        "Sub": "-"
    },
    "Loop": {
        "Start": "[",
        "End": "]"
    }
}
</code></pre>

#### `main.sd` example 

<pre><code>+++++ 
[
    >+++++ +++++ +++
    <-
]
>
+.
-.
>+++
[
    <+++++
    >-
]
<++
.
++.
-----.
++++.
>++
[
    <+++
    >-
]
<+.
</code></pre>
> Prints: "BARTOSZ"

### Dependencies: 
You only need a <a href="https://golang.org/dl/">Golang compiler</a> to be able to run the program  

### Installation
After installing golang on your machine: 
* Clone this repository 
    * ``git clone https://github.com/BartoszBurgiel/snareDrum``
* Build the program 
    * ``go build -o snareDrum/backend/cmd/main.go <name-of-the-executable>``
* Have fun! 
    * ``./snareDrum.deb exec fibonacci.sdexe``
