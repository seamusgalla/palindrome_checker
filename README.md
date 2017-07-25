# Palindrome Checker

This program checks if a word or series of latin charachters is a palindrome i.e reads the same forwards as backwards. It is not case sensitive.  
This program was created for demo purposes.

## Usage

``ispal level``  
This outputs `true`

``ispal palindrome``  
This outputs `false`

Arguments can also be enclosed in single or double quotes.

### Valid Charachters
Currently the program only takes the basic latin alphabet:  
abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ  

An invalid input `ispal gr8` will give the following error `invaldChar:'8'`

### Multiple Arguments
Aruments should be split with a single space ' '  
``ispal level palindrome gr8``  
This will output `true false invalidChar:'8'`

## Options

### All Valid
If the `-allValid` flag is present the program will exit if any of the arguments are invalid.

``ispal -allValid level test invalid55``  
This will output `Argument 3 => invalidChar:'5'`  

### White Space and Puncuation Marks
The program can ignore white space or puncuation marks if desired.

``ispal -ignorePunc race,car``  
This will output `true`

To analyse arguements with white space the arguemetns must be enclosed by double or single quotes. Tabs and return charachters are considered white space.
``ispal -ignoreWhite "Was it a car or a cat I saw"``  
This will output `true`

### Custom Delimiter
The `-d` flag can be used to specify a custom delimiter between the output results. By default this is `' '`.  
e.g.
`ispal -d=", " news level kayak` => `false, true, true`

### Case Sensitive
The `-caseSen` flag makes the palindrome check case sensitive

e.g.
`ispal -caseSen "Kayak"` => `false`  
`ispal -caseSen "kayak"` => `true`

## Build from source
To build this package from source please use the latest version of go and git and have the [gb](https://getgb.io) package manager installed.

1. Navigate to your desired directory and clone the repo using  
`git clone https://github.com/seamusgalla/palindrome_checker.git`

2. Build the ispal tool using  `gb build all` from the same directory.
This will create a the executable `bin\ispal.exe`

You can now use the program e.g. `bin\ispal kayak` => `true`

## Licence

MIT
