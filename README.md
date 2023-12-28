
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/iudicium/pryingdeep/web/static/logo.png">
    <img src="web/static/logo.png" alt="Logo" width="350" height="200">
  </a>

  <h3 align="center">Prying Deep</h3>

  <p align="center">
    <b>OSINT intellegence gathering tool for the dark/clear web
    <br />
    <br />
    <a href="https://github.com/iudicium/pryingdeep/issues">Report Bug</a>
    <a href="https://github.com/iudicium/pryingdeep/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#manual-installation">Manual Installation</a></li>
        <li><a href="#binary-installation">Binary Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#docker">Docker</a>
      <ul>
        <li><a href="#tor">Tor</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#preview">Preview</a></li>
  </ol>
</details>



## Translation (Перевод)
 - [Click](./docs/RU_README.MD)

<!-- ABOUT THE PROJECT -->
## About The Project

I've seen multiple projects out there in GitHub, that are crawlers for the deep web,
but most of them did not meet *my* standards of OSINT on the deep web.
So I decided to create my own deep web OSINT tool.


This tool serves as a reminder that the best practices of OPSEC should always be followed in the deep web.

<u>The author of this project is not responsible for any possible harm caused by the usage of this tool.</u>

## What does this tool do?

*Prying Deep crawls dark/clear net platforms and extracts as much intelligence as possible effectively.*


---

<!-- GETTING STARTED -->
## Getting Started


### Prerequisites

Before you can use our OSINT tool, please ensure you have the following dependencies installed:

1. **Docker: (optional)**
    - You can download and install Docker by following the official installation instructions for your specific operating system:
    - [Docker Installation Guide](https://docs.docker.com/get-docker/).

2. **Go: (required)**
    - [Golang Installation Guide](https://go.dev/doc/install).

3. **PostgresSQL: (required if you don't use docker)**
    - Make sure your `pryingdeep.yaml` environment variables match the environment in `docker-compose.yaml`
    - [PostgreSQL Installation](https://www.postgresql.org/download/)


### Binary Installation

1. Install the binary via:

```Sh
go install -v github.com/iudicium/pryingdeep/cmd/pryingdeep@latest
```

2. Run the following command:
```sh
pryingdeep install
```
3. Adjust the values inside the config folder to your needs.

### Manual Installation

1. Clone the repo: 
   ```sh
    git clone https://github.com/iudicium/pryingdeep.git
   ```

2. Adjust the values in the .yaml configuration either through flags or manually.
   `Database`, `logger`, `tor` all require manual configuration. <br>

<u>You *will need to read*</u> [Colly Docs](https://github.com/gocolly/colly/blob/v1.2.0/colly.go). Also, refer to [Config Explanation](./docs/CONFIG.MD#table-of-contents)


3. Build the binary via:
   - `go build` -> inside the `cmd/pryingdeep` directory
   - `go build cmd/pryingdeep/pryingdeep.go` -> root directory, binary will also be there.

## Docker

To start run *pryingdeep* inside a docker container use this command:
```sh
docker-compose up
```


## Config
Read more each parameter here:
[config](./docs/CONFIG.MD#table-of-contents)


## Tor
Read more about building and running our tor container here:
[Tor](./build/package/tor/README.MD#build)


<!-- USAGE EXAMPLES -->
## Usage

```
Pryingdeep specializes in collecting information about dark-web/clearnet websites.
This tool was specifically built to extract as much information as possible from a .onion website

Usage:
pryingdeep [command]

Available Commands:
completion  Generate the autocompletion script for the specified shell
crawl       Start the crawling process
export      Export the collected data into a file.
help        Help about any command
install     Installation of config files

Flags:
-c, --config string   Path to the .yaml configuration. (default "pryingdeep.yaml")
-h, --help            help for pryingdeep
-z, --save-config     Save chosen options to your .yaml configuration
-s, --silent          -s to disable logging and run silently


```


<!-- ROADMAP -->
## Roadmap

- [x] Add a search command which will not require any onion links
- [ ] Acquire a shodan api key for testing the favicon module
- [ ] Think of a way to acquire IP address of the server
- [ ] Implement scan command
- [ ] Implement file identification and search
- [x] Find  a way to supress gorm unique duplicate errors as they take half the screen...
<!-- CONTRIBUTING -->
## Contributing


If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request




<!-- LICENSE -->
## License

Distributed under the GPL-3.0 license. See `LICENSE` for more information.


## Video Preview

[![Video Title](https://img.youtube.com/vi/TvRT_lvixI0/maxresdefault.jpg)](https://www.youtube.com/watch?v=TvRT_lvixI0)

---

## Preview

![Preview](./web/static/preview.png)
![Data Preview](./web/static/data_preview.png)
![Data Preview 2](./web/static/data_preview2.png)


## Support

If you have found this repository useful and feel generous, you can donate some Monero (XMR) to the following address:

48bEkvkzP3W4SGKSJAkWx2V8s4axCKwpDFf7ZmwBawg5DBSq2imbcZVKNzMriukuPqjCyf2BSax1D3AktiUq5vWk1satWJt

Thank you!