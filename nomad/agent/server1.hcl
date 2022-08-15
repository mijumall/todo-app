data_dir = "/home/normal/space/data/"

region = "Tokyo"
datacenter = "Moon"

bind_addr = "192.168.155.98"

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

server {
	enabled = true
	bootstrap_expect = 1
}

log_level = "INFO"

