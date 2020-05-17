## Compiler vs Interpreter

Unlike a compiler, an interpreter reads and runs the source code fed into it.
Usually an interpreter produces a half staged code, AST, (Abstract syntaxt tree), and then walks the AST and interprets it,
and these are called tree-walking interpreters.

Other than other small things, both share the same steps and functionality, like lexycal anylisys and parsing.

## Monkey

Monkey is the intrepreter I'll be building.

It will tokenize and parse Monkey code in a REPL (Read-eavl-print-loop) building up an internal representation of the code
as an AST and then evaluate it.

### How its composed

- the lexer
- the parser
- the Abstract Syntax Tree (AST)
- the internal object system
- the evaluator