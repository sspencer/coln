# COLN: Print Column

Print specified column from text file.  This replaces the [fish shell](https://fishshell.com/) 
function coln.fish which is slow for larger files.

```fish
function coln
    while read -l input
        echo $input | awk '{print $'$argv[1]'}'
    end
end
```

## Examples

```
$ cat input.txt
one two three
1 2 3

$ echo "Use with Pipes"
$ cat input.txt | coln 2
two
2 

$ echo "Use with Files"
$ coln 3 input.txt
three
3

$ echo "One flag: -q Strip quotes"
$ cat input2.txt
a = "one";
b = "two";

$ coln -q 3 input2.txt
one
two
```

## Install

Go makes it easy:

    go install
