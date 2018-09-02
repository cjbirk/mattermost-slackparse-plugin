plugin:
	go build -o plugin plugin.go
	tar -czvf plugin.tar.gz plugin plugin.yaml

clean:
	rm -f *.tar.gz
	rm -f plugin
