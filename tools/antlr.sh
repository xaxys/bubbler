#!/bin/bash

cur_dir=$(dirname "$0")
export CLASSPATH=".:$cur_dir/antlr-4.13.1-complete.jar:$CLASSPATH"
export TEST_CURRENT_DIR=${CLASSPATH//.:}
if [ "$TEST_CURRENT_DIR" == "$CLASSPATH" ]; then
    export CLASSPATH=".:$CLASSPATH"
fi
java org.antlr.v4.Tool "$@"
