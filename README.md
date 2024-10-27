# CCCalc: Another CLI Calculator

CCCalc is a command-line interface (CLI) calculator designed to efficiently parse and evaluate simple arithmetic expressions. This project is part of a coding challenge from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-calculator).

## Project Overview

CCCalc is built to handle a wide variety of arithmetic expressions involving basic mathematical operations. It's designed to be both a practical tool for quick calculations and a fun project to explore parsing algorithms and command-line utility development.

## Prerequisites

Before installing and running CCCalc, ensure you have the following installed:

- **Go (at least version 1.13)**: CCCalc is written in Go. You will need to have Go installed on your machine to compile and run the calculator. Download and install it from [Go's official website](https://golang.org/dl/).

## Features

- Evaluate arithmetic expressions including operations for addition, subtraction, multiplication, and division.
- Handle complex nested expressions with proper order of operations.
- Error handling for common mistakes in arithmetic expressions, such as division by zero.

## Installation

Follow these steps to install CCCalc:

```bash
git clone https://github.com/nullsploit01/cc-calculator.git
cd cc-calculator
go build -o cccalc
```

## Usage

To use cccalc, you can run the compiled binary directly from the command line:

```bash
./cccalc "expression"
```

## Examples

```bash
./cccalc "3 + 4 * 2 / ( 1 - 5 )"
./cccalc "5 + 2 * (3 - 1)"
./cccalc "10 / 2 + 3 * (2 + 1)"
```
