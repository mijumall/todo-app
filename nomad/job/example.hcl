job "example" {
  region = "Tokyo"
  datacenters = ["Moon"]

  group "cache" {
    count = 2

    volume "sample-v" {
      type = "host"
      source = "sample-v"
    }

    task "cache" {
      driver = "docker"
    
      config {
        image = "redis:7"
        auth_soft_fail = true
      }
    
      volume_mount {
        volume = "sample-v"
        destination = "/v"
      }

      resources {
        cpu = 500
        memory = 256
      }
    }
  }
}
