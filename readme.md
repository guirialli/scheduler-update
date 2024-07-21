# Scheduler Update

Scheduler Update is a tool for Linux systems that automates the updating of apt and flatpak packages through scheduling. This project allows you to keep your system consistently updated efficiently.

## Installation

To start using Scheduler Update, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/guirialli/scheduler-update.git
   ```

2. **Compile the code:**

   Navigate to the project directory and build using `go build`:

   ```bash
   cd scheduler-update
   go build ./cmd/update/scheduler_update.go
   ```

   This will create an executable named `scheduler_update`.

3. **Move the compiled executable to the desired directory (recommended `/opt/scheduler-update`):**

   ```bash
   sudo mv scheduler_update /opt/scheduler-update
   ```

## Setting Up the Service

To run Scheduler Update as a systemd service at the system level, follow these steps:

1. **Create a systemd service file:**

   Create a file named `scheduler-update.service` in `/etc/systemd/system/` with the following content:

   ```ini
   [Unit]
   Description=Scheduler Update Service
   After=network.target

   [Service]
   Type=simple
   ExecStart=/opt/scheduler-update/scheduler_update
   Restart=always

   [Install]
   WantedBy=multi-user.target
   ```

2. **Enable and start the service:**

   Execute the following commands to enable and start the service:

   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable scheduler-update.service
   sudo systemctl start scheduler-update.service
   ```

Scheduler Update is now configured to run automatically according to your schedule, keeping your apt and flatpak packages updated.

---