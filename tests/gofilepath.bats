#!/usr/bin/env bats

@test "ext" {
  run gofilepath ext 'gofilepath.bats'
  [ "$output" = ".bats" ]
}

@test "base" {
  run gofilepath base 'gofilepath/main.go'
  [ "$output" = "main.go" ]
}

@test "dir" {
  run gofilepath dir 'a/b/c'
  [ "$output" = "a/b" ]
}

@test "split" {
  run gofilepath split "a/b.go" -o json
  [ "$output" = '{"dir":"a/","file":"b.go"}' ]

  run gofilepath split "a/b.go"
  [ "$output" = "$(printf 'a/\nb.go\n')" ]
}
