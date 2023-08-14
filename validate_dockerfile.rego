package main

import future.keywords
# import Dockerfile as input


# Use multi-stage builds
default multi_stage := false
multi_stage := true {
    input[_].Cmd == "copy"
    val := concat(" ", input[_].Flags)
    contains(lower(val), "--from=")
}
deny[msg] {
    multi_stage == false
    msg := sprintf("Not a multi-stage Dockerfile", [])
}


# Disallows latest tag in images 
disallow_latest_tag[msg]{
    input[_].Cmd == "from"
    tag:= split(input[_].Value[0],":")
    contains(lower(tag[0]),"latest")

    msg:=sprintf("Use of latest tag is Dissallowed",[])
}

# Disallow root user within container
root_users:=["root","0"]

deny_root[msg] {
    command := "user"
    users := [name | input[_].Cmd == "user"; name := input[_].Value]
    user := users[count(users)-1]
    contains(lower(lastuser[_]), forbidden_users[_])
    msg = sprintf(" USER %s is disallowed", [user])
}