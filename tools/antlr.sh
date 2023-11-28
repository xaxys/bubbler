#!/bin/bash

cur_dir=$(dirname "$0")
CLASSPATH=".:$cur_dir/antlr-4.13.1-complete.jar:$CLASSPATH"
TEST_CURRENT_DIR=${CLASSPATH//.:}
if [ "$TEST_CURRENT_DIR" == "$CLASSPATH" ]; then
    CLASSPATH=".:$CLASSPATH"
fi
java org.antlr.v4.Tool "$@"
