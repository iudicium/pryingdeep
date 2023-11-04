
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/pryingbytez/pryingdeep/web/static/logo.png">
    <img src="web/static/logo.png" alt="Logo" width="350" height="200">
  </a>

  <h3 align="center">Prying Deep</h3>

  <p align="center">
    <b>An OSINT intellegence gathering tool for the dark web
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/pryingbytez/pryingdeep/issues">Report Bug</a>
    <a href="https://github.com/pryingbytez/pryingdeep/issues">Request Feature</a>
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
        <li><a href="#installation">Installation</a></li>
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
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

I've seen multiple projects out there in Github, that are crawlers for the deep web, but most of them did not meet *my* standards of OSINT on the deep web. So I decided to create my own deep web OSINT tool.


This tool serves as a reminder that best practices of OPSEC should always be followed in the deep web.

The author of this project is not responsible for any possible harm caused by the usage of this tool.


---
Note:
 ⚠️ *PhoneNumber* module is currently not working. Please do not use it as of now.
---

<!-- GETTING STARTED -->
## Getting Started


### Prerequisites

Before you can use the our osint tool, please ensure you have the following dependencies installed:

1. **Docker: (optional)**
    - You can download and install Docker by following the official installation instructions for your specific operating system: [Docker Installation Guide](https://docs.docker.com/get-docker/).

2. **Go: (required)**
    - [Golang Installation Guide](https://go.dev/doc/install).
3. **PostgresSQL: (required)**
    - You will also be able to start it from docker soon
    - [PostgreSQL Installation](https://www.postgresql.org/download/)

## Configuration 

| Config File          | Description                                                                                                                                                                                          | Command                          | AutoSave |
|----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|----------|
| `crawlerConfig.json` | You can configure things like user agents, rate limiting, and other crawl-related parameters. There will be a dedicated section for the parameters later. Please refer to the gocolly docs for now.) | Manual Configuration             | No       |
| `exporterConfig.json` | The exporter configuration file specifies how data collected by the crawler is exported or saved. You can define the output format, file paths, and export options here.                             | pryingdeep&nbsp;export&nbsp;json | Yes      |
| `pryingConfig.json`   | Configure the OSINT modules here. You can specify them all via the command line. It will also automatically run the crawler with the passed on settings                                              | pryingdeep&nbsp;crawl&nbsp;-h    | Yes      |

*Note:* `crawlerConfig.json` and `pryingConfig.json` are used for the `crawl` command.

### Installation

1. Clone the repo: 
   ```sh
    git clone https://github.com/pryingbytez/pryingdeep.git     
   ```

2. Adjust these settings to your needs.  They are located in `pryingdeep/configs/json`. <br>
<u>You *will need to read*</u> [Colly Docs](https://github.com/gocolly/colly/blob/v1.2.0/colly.go). Also, refer to [Crawler Configuration](#crawler-configuration)


3. Set the `.env` settings accoridng to your needs, the <u>SOCKS</U> is where tor's  port and host is running.

4. Build the binary via 
   - `go build` -> directory: `cmd/pryingdeep`
   - `go build cmd/pryingdeep/pryingdeep.go` -> root directory, binary will also be there. 

## Docker

To start run *pryingdeep* inside a docker container use this command:
```sh
docker-compose up
```
<u>Adjust</u> the sleep value inside of docker-compose.yaml! It depends on how fast your computer can start the tor container. 
Otherwise the program will panic.
```sh
 command:
      - '/bin/sh'
      - '-c'
      - '/bin/sleep 8(this value) && ./pryingdeep crawl '
```

## Tor
Read more about building and running our tor container here: 
[Tor](./build/package/tor/README.MD#build) 


<!-- USAGE EXAMPLES -->
## Usage

Crawler:
   ```sh
   ./pryingdeep crawl
   ```

Exporter:
 ```sh
 ./pryingdeep export
  ```



<!-- ROADMAP -->
## Roadmap

- [ ] Add a folder for storing configuration files so that go install could be used.
- [ ] Fix the phoneNumber module, get it to at least 50%
- [ ] Add email validation
- [ ] Add PostgreSQL to docker-compose

See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a full list of proposed features (and known issues).




<!-- CONTRIBUTING -->
## Contributing


If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the  GPL-3.0 license . See `LICENSE` for more information.
