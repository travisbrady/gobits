ping: ping.8
	8l -o ping ping.8

ping.8: ping.go
	8g ping.go

