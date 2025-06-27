# AutoPort

AutoPort is a simple, cross-platform tool for automatic port forwarding using UPnP/NAT-PMP. It is designed to help you and your friends easily play games like Minecraft together, even if you are behind a router or firewall.

## Features
- Automatic port forwarding (UPnP/NAT-PMP)
- Cross-platform: Linux, macOS, Windows
- Easy to use, no configuration required

## Disclaimer
> **Note:** AutoPort will not work on all networks. Some routers or ISPs may block UPnP/NAT-PMP or restrict port mapping. Use at your own risk.

## How to Build

You need [Go](https://golang.org/dl/) installed.

```
# Clone the repository (after you git init and push to GitHub)
git clone https://github.com/DrFineSir/AutoPort.git
cd AutoPort

# Build for your platform
make build

# Or build for a specific platform
GOOS=linux GOARCH=amd64 go build -o AutoPort-linux main.go
GOOS=darwin GOARCH=amd64 go build -o AutoPort-mac main.go
GOOS=windows GOARCH=amd64 go build -o AutoPort-windows.exe main.go
```

## How to Run

```
# On Linux/macOS
./AutoPort-linux

# On Windows
AutoPort-windows.exe
```

By default, AutoPort will attempt to forward a random port to your local Minecraft server (default port 25565). Share your external IP and mapped port with your friends so they can join your game.

## Why use AutoPort?
- No more complicated router setup
- Play Minecraft or other games with friends easily
- Open source and free

## Troubleshooting
- Make sure UPnP or NAT-PMP is enabled on your router
- Some networks (e.g., campus, corporate, or CGNAT) may block port mapping
- Check your firewall settings

## License
MIT
