# quote-collect

Collect and organize quotes in the terminal. Store, search, and discover quotes by category, author, or favorites.

## Features

- Add quotes with author, source, and category
- List quotes filtered by category, sorted by author/date/random
- Show a random quote for inspiration
- Mark quotes as favorites
- Search quotes by text, author, or source
- Export collection as JSON
- Statistics (total, favorites, breakdown by category)
- SQLite persistence in `~/.quote-collect/quotes.db`

## Install

```bash
go install github.com/TataneSan/quote-collect@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/quote-collect.git
cd quote-collect
go build -o quote-collect .
sudo mv quote-collect /usr/local/bin/
```

## Usage

```
quote-collect add <text> [--author NAME] [--source SOURCE] [--category CAT]
quote-collect list [--category CAT] [--sort id|author|date|random] [--fav]
quote-collect show <id>
quote-collect random
quote-collect fav <id>
quote-collect remove <id>
quote-collect search <query>
quote-collect stats
quote-collect export [--json]
```

## Examples

```bash
# Add quotes
quote-collect add "The only way to do great work is to love what you do" --author "Steve Jobs" --category inspiration
quote-collect add "Talk is cheap. Show me the code." --author "Linus Torvalds" --category programming
quote-collect add "Simplicity is the ultimate sophistication." --author "Leonardo da Vinci" --category design

# List all
quote-collect list

# List favorites
quote-collect list --fav

# Random quote
quote-collect random

# Toggle favorite
quote-collect fav 1

# Search
quote-collect search code

# Stats
quote-collect stats

# Export
quote-collect export --json
```

## License

MIT
