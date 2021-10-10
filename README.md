<div id="top"></div>

<br />
<div align="center">

  <h3 align="center">Goministrator</h3>

  <p align="center">
  A Discord administration and assistant bot for your server.
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Goministrator is a Discord management, administration and server assistant. Which does the following:
* Setup meetings in a server
* Summarise and create audio and text transcripts of meetings
* Email transcripts to meeting participants



<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Golang](https://golang.org/) 
* [DiscordGo](https://github.com/bwmarrin/discordgo/)

<p align="right">(<a href="#top">back to top</a>)</p>


## Getting Started

### Prerequisites

* [Golang](https://golang.org/doc/install) 

### Installation

1. Get a free API Key at [https://discord.com/developers](https://discord.com/developers)
2. Clone the repo
   ```sh
   git clone https://github.com/TheLazyLemur/Goministrator.git
   ```
3. Install Goministrator
   ```sh
   go install .
   ```

<p align="right">(<a href="#top">back to top</a>)</p>


## Usage

Run Bot:
```bash
Goministrator -bot -token <DiscordToken>
```

<p align="right">(<a href="#top">back to top</a>)</p>


## Roadmap

- [] Audio Transcript
    - [x] Record and save audio
    - [] Transcribe audio to text
    - [] Summarise text
    - [] Email Summarised text, full text and audio to participants
- [] Channel administration
    - [] Create meetings
    - [] Create channel before meeting
    - [] Create roles for restricitng meeting participation


<p align="right">(<a href="#top">back to top</a>)</p>


## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [DiscordGo](https://choosealicense.com)

<p align="right">(<a href="#top">back to top</a>)</p>
