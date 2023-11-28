#!/bin/bash
cur_dir=$(dirname "$0")
"$cur_dir/antlr.sh" -Dlanguage=Go -package parser -o "$cur_dir/../parser" -no-listener -visitor "$cur_dir/../bubbler.g4"