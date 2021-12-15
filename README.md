# COLN: Print Column

Print specified column from text file.

```
$ cat input.txt
one two three
1 2 3

$ cat input.txt | coln 2
two
2 

$ coln 3 input.txt
three
3

$ cat input2.txt

```

## Install

Go makes it easy:

    go install
