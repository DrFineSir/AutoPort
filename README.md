# AutoMap

AutoMap is a simple, cross-platform tool for automatic port forwarding using UPnP/NAT-PMP. It is designed to help you and your friends easily play games like Minecraft together, even if you are behind a router or firewall.

## Features
- Automatic port forwarding (UPnP/NAT-PMP)
- Cross-platform: Linux, macOS, Windows
- Easy to use, no configuration required

## Disclaimer
> **Note:** AutoMap will not work on all networks. Some routers or ISPs may block UPnP/NAT-PMP or restrict port mapping. Use at your own risk.

## How to Build

You need [Go](https://golang.org/dl/) installed.

```
# Clone the repository (after you git init and push to GitHub)
git clone https://github.com/yourusername/AutoMap.git
cd AutoMap

# Build for your platform
make build

# Or build for a specific platform
GOOS=linux GOARCH=amd64 go build -o AutoMap-linux main.go
GOOS=darwin GOARCH=amd64 go build -o AutoMap-mac main.go
GOOS=windows GOARCH=amd64 go build -o AutoMap-windows.exe main.go
```

## How to Run

```
# On Linux/macOS
./AutoMap-linux

# On Windows
autoMap-windows.exe
```

By default, AutoMap will attempt to forward a random port to your local Minecraft server (default port 25565). Share your external IP and mapped port with your friends so they can join your game.

## Why use AutoMap?
- No more complicated router setup
- Play Minecraft or other games with friends easily
- Open source and free

## Troubleshooting
- Make sure UPnP or NAT-PMP is enabled on your router
- Some networks (e.g., campus, corporate, or CGNAT) may block port mapping
- Check your firewall settings

## License
MIT

