# Not tried yet.
# Run this job after running client agent as root

job "example" {
  region = "Tokyo"
  datacenters = ["Moon"]

  group "cache" {
    count = 2

    network {
      mode = "bridge"
      port p1 {}
    }

    volume "sample-v" {
      type = "host"
      source = "v"
    }

    task "python-server" {
      driver = "docker"
    
      config {
        image = "python"
        auth_soft_fail = true

        command = "sleep"
        args = ["100000"]
      }
    
      volume_mount {
        volume = "sample-v"
        destination = "/v"
      }

      env {
        IP_ADDR = "${attr.unique.network.ip-address}"
        GREETINGS = "Hello under world"
      }

      resources {
        cpu = 500
        memory = 256
      }
    }
  }
}
