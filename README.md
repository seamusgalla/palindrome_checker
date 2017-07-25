# Palindrome Checker

This program checks if a word or series of latin charachters is a palindrome i.e reads the same forwards as backwards. It is not case sensitive.  
This program was created for demo purposes.

## Usage

``ispal level``  
This outputs `true`

``ispal palindrome``  
This outputs `false`

Arguments can also be enclosed in single or double quotes.

### Invalid Charachters

``ispal gr8``  
This will give the following error `invaldChar:'8'`

### Multiple Arguments
``ispal level palindrome gr8``  
This will output `true false invalidChar:'8'`

### Options

#### All Valid
If the `-allValid` flag is present the program will exit if any of the arguments are invalid.

``ispal -allValid level test invalid55``  
This will output `Argument 3 => invalidChar:'5'`  

#### White Space and Puncuation Marks
The program can ignore white space or puncuation marks if desired.

``ispal -ignorePunc race,car``  
This will output `true`

To analyse arguements with white space the arguemetns must be enclosed by double or single quotes. Tabs and return charachters are considered white space.
``ispal -ignoreWhite "Was it a car or a cat I saw"``

## Licence

MIT
