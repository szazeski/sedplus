#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
TEST_COUNT=0

go build

function sedtest() {
  TEST_COUNT=$((TEST_COUNT+1))
  echo "$TEST_COUNT) $3 -> $2"
if [ "$1" != "$2" ]; then
    echo "   FAIL: expected '$2'"
    echo "           actual '$1'"
    exit 1
else
  echo "   PASS"
fi
}

actual=$(echo "Hello World" | ./sedplus)
sedtest "$actual" "Hello World" "should pass through input"

actual=$(echo "Hello World" | ./sedplus --find 'Hello' --replace 'Goodbye')
sedtest "$actual" "Goodbye World" "piping Hello World"

actual=$(echo "Hello World\nHello Second Line" | ./sedplus --find 'Hello' --replace 'Goodbye')
sedtest "$actual" "Goodbye World\nGoodbye Second Line" "checking multiline replace"

#actual=$(echo "Hello World\nHello Second Line" | ./sedplus --find-line 'Second' --replace 'This is now the third line')
#sedtest "$actual" "Hello World\nThis is now the third line" "checking full line replace"
# newline doesn't seem to work here

actual=$(echo "Hello World" | ./sedplus --uppercase)
sedtest "$actual" "HELLO WORLD" "--uppercase"

actual=$(echo "Hello World" | ./sedplus --lowercase)
sedtest "$actual" "hello world" "--lowercase"

actual=$(echo "     Hello     World     " | ./sedplus --trim)
sedtest "$actual" "Hello     World" "--trim piping '     Hello      World     '"

actual=$(echo "Hello 1234 World !*7" | ./sedplus --numeric)
sedtest "$actual" "12347" "--numeric"

actual=$(echo "Hello 1234 World !*7" | ./sedplus --alpha)
sedtest "$actual" "HelloWorld" "--alpha"

actual=$(echo "Hello 1234 World !*7" | ./sedplus --alphanumeric)
sedtest "$actual" "Hello1234World7" "--alphanumeric"


echo "${GREEN}All Tests PASS${NC}"
