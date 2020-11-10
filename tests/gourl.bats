#!/usr/bin/env bats

@test "pathescape" {
  [ "$(gourl pathescape 'a+b+/c' | gourl pathunescape)" = "a+b+/c" ]
}

@test "parse" {
  [ "$(gourl parse 'https://golang.org/pkg/net/url/#PathEscape' | jq -r .Host)" = "golang.org" ]
}
