.PHONY: install build clean

build: ../logviewer

../logviewer:
	go build

install: ../logviewer
	sudo mkdir /usr/local/logviewer
	sudo cp logviewer /usr/local/logviewer/logviewer
	sudo cp -r index.html /usr/local/logviewer/
	sudo cp -r fonts /usr/local/logviewer/
	sudo cp -r css /usr/local/logviewer/
	sudo cp -r js /usr/local/logviewer/
	sudo cp ./init-script /etc/init.d/logviewer

clean:
	rm ../logviewer
	sudo rm -rf /usr/local/logviewer
	sudo rm -rf /etc/init.d/logviewer
