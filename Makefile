ping: ping.8
	8l -o ping ping.8

ping.8: ping.go
	8g ping.go

clean:
	rm ping.8 ping	

.PHONY : clean
