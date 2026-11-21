---
layout: default
title: Language notes
---

## Expressions the evaluator can run

```text
2 * (3 + 4)
10 / 2
3 < 8
!false
-12
```

The current evaluator supports integer `+`, `-`, `*`, `/`, `<`, `>`, `==`, and `!=`, boolean literals, `!`, and unary `-`.

These expressions evaluate to `14`, `5`, `true`, `true`, and `-12`, respectively.

## Still being implemented

The lexer and parser also recognize identifiers, `let`, `return`, `if`/`else`, function literals, and calls. The evaluator does not implement these constructs yet, so successfully parsing them does not mean they can run. Unsupported statements are reported by the playground rather than producing a value.
