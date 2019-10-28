# Syllabus
Idiotic script to check the availability of all the consonnant-starting one-syllabus domain name, with the .com and .io TLD, using the `whois` CLI program.
It's super slow (no parallelism) to play fair with the whois API


## Usage
```
go cmd/main.go [word] > results.csv
```
Where all the names before 'word' are ignored (crude journaling).
The result is a CSV where each row is of the shape `<name>,<is .com available>,<is .io available>`
