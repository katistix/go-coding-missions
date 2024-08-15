# Hacker News Time Machine

## The Scenario 🕰️

You've just invented a time machine (congratulations!) but there's a catch - it runs on tech news. To fuel your journey back to 2050, you need to scrape the hottest headlines from Hacker News circa 2024. But be careful, you don't want to disrupt the space-time continuum!

## Your Mission 🚀

Create a Go program that:

1. Scrapes the top 50 post titles from Hacker News (https://news.ycombinator.com/)
2. Uses goroutines and channels because... future tech is all about concurrency!
3. Implements a "quantum flux capacitor" (fancy talk for rate limiting) to avoid detection
4. Handles any hiccups along the way (errors are just temporal anomalies, right?)
5. Outputs the titles in the same order they appear on the page (maintaining the timeline is crucial!)

## Rules of Time Travel 📋

- Use Go (obviously, it's the language of the future)
- No external libraries for scraping (sorry, that's cheating across timelines)
- You can use `golang.org/x/net/html` for parsing HTML (even time travelers need some help)

## Example Output 🖥️

"Researchers discover new quantum algorithm"

"startup.io raises $50M for AI-powered toasters"

"The hidden dangers of time travel: a developer's perspective"

...

## Bonus Points 🌟

- Implement a caching mechanism (temporal storage unit)
- Add a command-line flag to specify the number of posts to fetch
- Calculate and display the time taken to complete the operation (in multiple timelines, of course)

## Hint 💡

Remember, young time traveler, concurrency is your friend. Use it wisely, and you'll be back to 2050 before you know it!

Happy coding, and safe travels! 🚀🕰️