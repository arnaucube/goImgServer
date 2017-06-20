# goImgCDN
CDN for images, written in Go lang

This is a server to post images and serve images.
File type accepted:
```
- PNG
- JPG
- JPEG
```
All the files are converted to PNG, and resized to the configured sizes.

## Use
upload example (with curl):
```
curl -F file=@./image1.png http://127.0.0.1:3050/image
```

to get image:
```
http://127.0.0.1:3050/images/image1.png
```

## Configuration
Example configuration file (config.json):
```
{
    "folder": "files",
    "blockedIPs": [
        "192.168.1.3",
        "147.116.48.158"
    ],
    "allowedIPs": [
        "127.0.0.1"
    ],
    "imgWidth": 200,
    "imgHeigh": 0
}
```
The "allowedIPs" are the IPs allowed to post images.

"blockedIPs" are the IPs blocked from posting images.

Also, all IPs not present in the "allowedIPs", will be blocked.
