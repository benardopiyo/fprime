# fprime

## Instructions

* The `fprime` program takes a single positive integer as an argument and displays its prime factors in ascending order, separated by a *, followed by a newline ('\n').


## Requirements

* The program display the prime factors in ascending order, separated by *.
* If the number of arguments is different from 1, or if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

## Usage

* To run the program, use the following commands in your terminal:

```bash
$ go run . 225225
3*3*5*5*7*11*13
```
```bash
$ go run . 8333325
3*3*5*5*7*11*13*37
```
```bash
$ go run . 9539
9539
```
```bash
$ go run . 804577
804577
```
```bash
$ go run . 42
2*3*7
```
```bash
$ go run . a

```
```bash
$ go run . 0

```
```bash
$ go run . 1

```