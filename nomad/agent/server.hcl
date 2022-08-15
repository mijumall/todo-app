data_dir = "/home/normal/space/data/"

region = "Tokyo"
datacenter = "Moon"

bind_addr = "0.0.0.0"

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

