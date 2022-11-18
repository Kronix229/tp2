#!/usr/bin/env bash

set -eu

PROGRAMA="$1"

RET=0
OUT=`mktemp`
trap "rm -f $OUT" EXIT

for x in *.test; do
  b=${x%.test}
  echo -n "Prueba $b... "
  cat "$x"
  echo

  ($PROGRAMA ${b}_usuarios <${b}_in || RET=$?) |
    diff -u --label "${b}_cátedra" --label "${b}_estudiante" ${b}_out - >$OUT || :

  if [[ $RET -ne 0 ]]; then
    echo -e "programa abortó con código $RET."
    

  elif [[ -s $OUT ]]; then
    echo -e "output incorrecto:\n"
    cat $OUT
    wait
    

  else
    echo -e "OK."
  fi
  echo
wait
done
