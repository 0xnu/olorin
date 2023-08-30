## Telegram Music Sharing Bot

[Olorin](https://www.wordhippo.com/what-is/the-meaning-of/yoruba-word-c156a5b85a427c85b7e23dba1c2d8187b7860ad7.html) is an advanced Telegram music bot written in Golang. It allows users (or admins/moderators) to share songs with dynamic emoji reactions and a Twitter profile link for further engagement.

### Key Features

Below are the key features of the bot:

* Music Sharing: Users can share songs in MP3 format. Each shared song has essential metadata like the title, artist, and genre.
* Dynamic Emoji Reactions: The bot provides a set of emojis (👍, 👎, 🔥, 🥰, 🚀) that users can tap to react to the shared song. The count of each emoji reaction updates in real time.
* User Choice Memory: The bot remembers the last emoji each user tapped. If a user changes their reaction, the bot updates the counts accordingly, ensuring that each user's reaction is unique at any given time.
* Twitter Profile Link: Each shared song includes A Twitter profile link. It allows users to follow the source or artist for more updates. The link is formatted to prevent Telegram from showing a web preview card, ensuring a clean and focused user interface.
* Command Deletion: When a user shares song using the `/sharemusic` command, the bot deletes the command message to keep the chat clutter-free.
* Concurrency-Safe: The application uses mutex locks to ensure that the emoji counts are updated atomically, making it safe for use in high-traffic groups.
* Debugging Enabled: The bot runs in debug mode, logging essential information for monitoring and troubleshooting.

### Instructions

- Grab a key from [Bot Father](https://telegram.me/BotFather).
- Update `BOT_TOKEN` and song details in [config.json](./config.json)

### Install Dependencies

First, you'll need to install the Telegram Bot API package for Golang:

```shell
go mod init main
go mod tidy
```

### Local Development

Fire up the bot:

```shell
go run .
```

### Deployment to Production

Alternatively, you can deploy your own copy of the app using the web-based flow:

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Screenshots

Screenshots of the bot are available [here](./screenshots/).

### Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/moatsystems/imdb_scrapy/tags).

### License

This project is licensed under the [Unlicense License](LICENSE) - see the file for details.

### Copyright

(c) 2023 [Finbarrs Oketunji](https://finbarrs.eu).

[![CC0](http://mirrors.creativecommons.org/presskit/buttons/88x31/svg/cc-zero.svg)](LICENSE)
