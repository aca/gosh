#!/usr/bin/env bats

@test "lookupaddr" {
  run gonet lookupaddr "8.8.8.8"
  [ "$output" = 'dns.google.' ]
}

@test "lookuphost" {
  run gonet lookuphost "google.com" -o json
  [ "$output" = '["172.217.175.238","2404:6800:4004:811::200e"]' ]
}

@test "lookuptxt" {
  run gonet lookuptxt "google.com" -o json
  [ "$status" -eq 0 ]
}
