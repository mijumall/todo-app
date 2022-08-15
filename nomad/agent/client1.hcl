data_dir = "/home/normal/space/data/"

region = "Tokyo"
datacenter = "Moon"

bind_addr = "192.168.189.230"

addresses {
	http = "localhost"
}

advertise {
	http = "localhost"
}

ports {
	http = 4646
	rpc = 4647
	serf = 4648
}

client {
	enabled = true
	servers = ["192.168.155.98"]

	host_volume "sample-v" {
		path = "/home/normal/space/v"
	}
}

log_level = "INFO"
