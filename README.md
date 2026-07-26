# go_git_contribution

A github-styled contribution heatmap for your local git repositories, rendered in the terminal. Built in go(2024).

Github's contribution graph only shows what you push to GitHub. This tool scans git repos on your machine - work repos, mirrors, side projects on other remotes - and renders the same familiar graph from your local commit history.

## How it works

1. Register folders to scan (stored in a dotfile in your home directory). The canner walks each foler recursively and records every `.git` directory it finds.
2. Render the graph for a given author email: the tool walks each repo's commit log, buckects commits by day for the last ~6 months, and prints a colored week-by-week grid with month labels - today highlighted.

## Usage
```
go build -o gitgraph .

# 1. regitster one or more folders containing git repos
./gitgraph -add ~/workspace

# 2. render your contribution graph
./gitgraph -email your@example.com
```

## Implementation notes

- Recursive repo discover with an ignore for `vendor/node_modules`-style noise
- Commit bucketing by calendar day with correct week-offset math so the grid aligns like Github's 
- Cell coloring by commit cound thresholds. Current day higlighted
- No external services - reads local git data only
