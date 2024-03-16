# Tech Radar

Based on the [Zalando Tech-Radar](https://github.com/zalando/tech-radar).

Available at <http://maroda.github.io/radar>.

This is my personal radar, which I have used as a framework for developing them elsewhere.

Since this is being hosted by Github Pages, it uses the optional `/docs/` directory to store the content.

## Local Dev

The root directory (with this README) can use `yarn` to serve the page for local development (it is configured to serve `/docs/`):

1. `brew install node yarn`
2. `yarn`
3. `yarn start`

The Radar will be available at <http://localhost:3000>.

## The `radar` processor

Build: `go build`
Usage: `./radar -list <structured_data>.csv`

This takes a CSV and spits out json for use with the radar.

The contents are for an array inside `docs/config.json`. See the discussion below for setting it up.

## The CSV file

The **Subject** column is whatever name you want to use.

The **Category** column matters, it should be one of four things:
1. Code
2. Data
3. Platform
4. Tool

The **Ring** column matters, it should be one of four things:
1. Adopt
2. Trial
3. Assess
4. Hold

The **Moved** column is not automated.

Example CSV:
```
Subject,Cat,Ring,Moved
1Password,Platform,Adopt,0
ArgoCD,Code,Adopt,0
Elixir,Code,Hold,0
```

## Full ETL Command

This was originally written before a config file was used for the radar. **TODO:** Make the `radar` command write out the entire `config.json`.

Right now it does not, so there's a skeleton framework:

First translate the CSV, change it all to lowercase, and chop off that nasty final comma:
```
./radar -list WMTRad-v1.csv | awk '{print tolower($0)}' | sed '$ s/.$//' > radout.json
```

Next cat the head and tail of the array on each end:
```
cat head.txt radout.json tail.txt > docs/config.json
```

That should get it all set, the static Tech Radar is in place and ready to view.

