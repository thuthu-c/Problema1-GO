# Parallel Array Summarizer Go

This program aims to obtain some specific information of a
generated array using multithreading.

## Running the program

On project root run the shell script with the positional arguments, like this

`$ ./start.sh <n> <t>`

Where:
- `n` means that the program will load the array with 10^n items;
- `t` is the number of threads that the program will create to process the array.

So, for instance, if you want to load the array with 10^7 items and need to use 32 threads, run:

```sh
$ ./start.sh 7 32
```

## Maintainers
- [Pedro Costa](mailto:pedroc_aragao@outlook.com)
- Thuanny Albuquerque
- Esther Maria
