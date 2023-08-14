package main

import future.keywords

import test-dockerfile as schema



# Use multi-stage builds
multi_stage[msg]{
   schema[_].Cmd == "from"
   schema[_].Stage == 0

   msg:=sprintf("Multi-stage Dockerfiles allowed")
}

disallow_latest_tag[msg]{
    schema[_].Cmd == "from"
    tag:= split (schema[_].Value[0],":")
    contains(lower(tag[0]),"latest")

    msg:=sprintf("Use of latest tag is Disallowed")
}

