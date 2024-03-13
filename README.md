# Tech Radar

Based on the [Zalando Tech-Radar](https://github.com/zalando/tech-radar).

Available here: <http://techradar>, contact me for the passphrase.

This repo has Pages enabled, SSL is required. It is only accessible via passphrase, there is no username, and the URL hash expires after one second.

## How It Works

The included JavaScript prompts for a password. The password is hashed and matched against a directory listing. If the listing returns HTTP code 200, the user is allowed access. If no directory with that hash exists, no 200 code is served, and the user is redirected to root (`/`).

If a second passes between the time the passphrase is entered and the same URL is used again (i.e. a browser refresh, giving it to someone else), the path is rendered invalid and the user must return to the login page and enter the passphrase again. This is done by setting a variable at the login page and reading it from the protected page.

## Acknowledgements

* https://github.com/chrissy-dev/protected-github-pages
