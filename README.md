https://adventofcode.com

## Session cookie

1. Find session cookie on AoC site after logging in under Chrome Developer
   Tools -> Application -> Cookies.
1. Create a file called `.session` with contents:

    ```shell
    session=YOUR_SESSION_COOKIE
    ```

1. Run the following to programmatically retrieve the day's input:

    ```shell
    curl -b "$(cat .session)" -o dayXX_input.txt https://adventofcode.com/20XX/day/XX/input
    ```
