# jsonstring

Convert stdin to a JSON-encoded string

# Installation

`go install github.com/racingmars/jsonstring`

# Example usage

Input:
```
% cat << EOF | jsonstring
this is my input string.
It contains newlines, and even
   a tab character!
Perhaps it also has "quoted text."
EOF
```

Output:
```
"this is my input string.\nIt contains newlines, and even\n\ta tab character!\nPerhaps it also has \"quoted text.\"\n"
```

