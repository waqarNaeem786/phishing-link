**For study purposes — use it on your own.**

https://github.com/user-attachments/assets/399cf93e-844b-436b-8ae3-a194de48cbed

## Overview

This is a simple Go application with a JavaScript frontend that extracts basic client information like IP, geolocation, platform, and browser language, then logs the data to the server.

## Features

* Captures user IP address
* Captures geolocation (latitude & longitude)
* Captures operating system/platform
* Captures browser language
* Logs the data to the terminal in a clean format

## Getting Started

1. Clone the repository:

```
git clone https://github.com/your-repo/client-info-collector.git
cd client-info-collector
```

2. Run the server:

```
go run main.go
```

3. Open `http://localhost:6969` in your browser to test.

## Deploying with ngrok

1. Download and install [ngrok](https://ngrok.com/).
2. Start the Go server:

```
go run main.go
```

3. Expose your port:

```
ngrok http 6969
```

Use the ngrok URL to access the tool from any device.

