#!/usr/bin/env bats

@test "trimspace" {
  run gostrings trimspace ' hello world '
  [ "$output" = "hello world" ]

  [ "$(echo -n ' hello world ' | gostrings trimspace)" = "hello world" ]
}

@test "trimspace multiline" {
  [ "$(printf ' hello world\nmulti ' | gostrings trimspace)" = "$(printf 'hello world\nmulti')" ]
}

@test "count" {
  run gostrings count 'cheese' 'ee'
  [ "$output" -eq "1" ]

  run gostrings count 'five' ''
  [ "$output" -eq "5" ]
}

@test "index 1" {
  run gostrings index 'chicken' 'ken'
  [ "$output" -eq 4 ]
}

@test "index 2" {
  run gostrings index 'chicken' 'dmr'
  [ "$output" -eq -1 ]
}

@test "hasprefix" {
  run gostrings hasprefix 'chicken' 'ch'
  [ "$status" -eq 0 ]

  run gostrings hasprefix 'chicken' '3ch'
  [ "$status" -eq 1 ]

  if echo -n 'chicken' | gostrings hasprefix 'ch'; then
    true
  else
    false
  fi

  if echo -n 'chicken' | gostrings hasprefix '3ch'; then
    false
  else
    true
  fi
}

@test "indexrune" {
  run gostrings indexrune 'chicken' 'k'
  [ "$output"  -eq 4 ]

  [ $(echo -n 'chicken' | gostrings indexrune 'k') -eq 4 ]
  [ $(echo -n 'chicken' | gostrings indexrune 'd') -eq -1 ]
}

@test "splitn" {
  run gostrings splitn 'a,b,c' ',' 2
  [ "$output" = "$(printf 'a\nb,c')" ]
}
