package main

import (
	"github.com/fatih/color"

	"github.com/pryingbytez/pryingdeep/pkg/cmd"
)

// TODO: add more options to save the config for crawlerConfig.Json
// TODO: add rate limiting.
// TODO:Fix phoneNumbers module, instead of extracting with regexp maybe just look for tel:
func main() {
	color.HiMagenta(art())
	cmd.Execute()
}

// view-source:http://xjfbpuj56rdazx4iolylxplbvyft2onuerjeimlcqwaihp3s6r4xebqd.onion/chatgpt-web-crawler/comment-page-1/
func art() string {
	return `
$$$$$$$\  $$$$$$$\$$\     $$\$$$$$$\ $$\   $$\  $$$$$$\  $$$$$$$\  $$$$$$$$\ $$$$$$$$\ $$$$$$$\
$$  __$$\ $$  __$$\$$\   $$  \_$$  _|$$$\  $$ |$$  __$$\ $$  __$$\ $$  _____|$$  _____|$$  __$$\
$$ |  $$ |$$ |  $$ \$$\ $$  /  $$ |  $$$$\ $$ |$$ /  \__|$$ |  $$ |$$ |      $$ |      $$ |  $$ |
$$$$$$$  |$$$$$$$  |\$$$$  /   $$ |  $$ $$\$$ |$$ |$$$$\ $$ |  $$ |$$$$$\    $$$$$\    $$$$$$$  |
$$  ____/ $$  __$$<  \$$  /    $$ |  $$ \$$$$ |$$ |\_$$ |$$ |  $$ |$$  __|   $$  __|   $$  ____/
$$ |      $$ |  $$ |  $$ |     $$ |  $$ |\$$$ |$$ |  $$ |$$ |  $$ |$$ |      $$ |      $$ |
$$ |      $$ |  $$ |  $$ |   $$$$$$\ $$ | \$$ |\$$$$$$  |$$$$$$$  |$$$$$$$$\ $$$$$$$$\ $$ |
\__|      \__|  \__|  \__|   \______|\__|  \__| \______/ \_______/ \________|\________|\__|
`
}
