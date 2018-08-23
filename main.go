package main

import (
    "gitee.com/johng/gf/g"
    _ "gitee.com/johng/gf-home/boot"
    _ "gitee.com/johng/gf-home/router"
)

func main() {
    s := g.Server("doc")
    s.SetDenyRoutes([]string{
        "/config/*",
        // "/static/template/*",
    })

    s.SetPort(9999)
    s.Run()
}