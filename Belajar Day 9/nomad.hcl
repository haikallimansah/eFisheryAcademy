job "haikallimansah-echo" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "http" {
        to = 1323
      }
    }

    task "haikallimansah-echo" {
      driver = "docker"

      config {
        image = "haikallimansah/go-echo:v1"
        ports = ["http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }

    service {
      name = "haikallimansah-echo"
      port = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.kowlon-echo-demo.rule=Host(\"haikallimansah.cupang.efishery.ai\")",
      ]
      check {
        port        = "http"
        type        = "tcp"
        interval    = "15s"
        timeout     = "14s"
      }
    }

  }
}