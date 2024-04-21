SRC_DIR := $(CURDIR)
DEST_DIR := /usr/local/src/simple-api-server-using-golang
SERVICE_NAME := simple-api-server-using-golang
EXECUTABLE := $(DEST_DIR)/$(SERVICE_NAME)
GO_BUILD := go build -o $(EXECUTABLE)
INSTALL_USER := $(shell whoami)

.PHONY: install uninstall start stop

install:
	# Compile Go code
	$(GO_BUILD) $(SRC_DIR)/main.go	
	# Copy service to systed
	sudo cp $(SRC_DIR)/$(SERVICE_NAME).service /etc/systemd/system/$(SERVICE_NAME).service
	# Copy needed files
	sudo cp $(SRC_DIR)/onlypepes.db $(DEST_DIR)
	sudo cp $(SRC_DIR)/.env $(DEST_DIR)
	# Set correct permissions for the executable file
	sudo chmod 755 $(EXECUTABLE)
	# Enable and start the service
	sudo systemctl enable $(SERVICE_NAME)
	sudo systemctl start $(SERVICE_NAME)

start:
	# Start the service
	sudo systemctl start $(SERVICE_NAME)

stop:
	# Stop the service
	sudo systemctl stop $(SERVICE_NAME)

uninstall:
	sudo systemctl stop $(SERVICE_NAME)
	sudo systemctl disable $(SERVICE_NAME)
	sudo rm /etc/systemd/system/$(SERVICE_NAME).service
	sudo rm -r $(DEST_DIR)