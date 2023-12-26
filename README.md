# GoJsonParser

GoJsonParser is an attempt to build a utility similar to JQ Parser from scratch. 

### **_NOTE_**:
Inspiration for this initiative was CodingChallenges.com : https://codingchallenges.fyi/challenges/challenge-json-parser/

## Supported Changes : 

- ✅ JSON Compiler : For every input string it does Lexical analysis and parses the tokens to predict if the input was valid JSON or not.
   1. Lexical Analysis : It reads the input JSON string byte by byte and convert into sequence of Tokens
   2. Parser : Post the tokens conversion from Lexical analysis, these tokens are evaluated against the JSON Grammar (https://www.json.org/json-en.html).
- ⬜️ JSON Query [TODO] : 

## Build 

```sh
go build .
```

## Run

```shell
go run .
```