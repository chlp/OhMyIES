# OhMyIES

**IES Student Telegram Notifier**

OhMyIES is a Go application that automatically connects to the **SchoolSoft** system, fetches new messages, categorizes them, and sends them to specific Telegram chats.

## Features

- Supports multiple SchoolSoft feeds
- Categorizes messages into:
  - **Lesson Planning**
  - **Absence**
  - **General**
- Sends categorized updates to corresponding Telegram chats via bots
- Simple configuration via JSON file
- Event logging to a file

## Installation

```bash
git clone https://github.com/chlp/OhMyIES.git
cd OhMyIES
go build -o ohmyies ./cmd/ohmyies
```

## Configuration

You need to create a `config.json` file based on the example provided.

```json
{
  "log_file": "app.log",
  "debug": true,
  "feeds": [
    {
      "name": "IES Vlad Absence",
      "key": "123",
      "key_2": "01d87e61af6df84253d5666634dc666",
      "chat": [
        {
          "type": "markdown",
          "bot_api_token": "YOUR_TELEGRAM_BOT_TOKEN",
          "chat_id": "YOUR_CHAT_ID"
        }
      ]
    }
  ]
}
```

**Fields:**
- `log_file` — path to the log file
- `debug` — enable debug output
- `feeds` — list of feeds to monitor
- `key` and `key_2` — access keys for the SchoolSoft feed
- `chat` — Telegram bot configuration for sending messages

## Usage

Run the application:

```bash
./ohmyies
```

The app will periodically check for new messages and send them to the corresponding Telegram chats.

## Requirements

- Go 1.21+
- Access to SchoolSoft
- A Telegram bot token

## License

MIT License
