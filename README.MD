# LeagueFixHyprland
### Simple tool to fix cursor for leagueoflegends-git on Hyprland with multi-monitor setup

## About
I experienced issues with cursor going outside the game after installation game with leagueoflegends-git as I had problems with Lutris (game does not start correctly). Wine's virtual environment did't work too.
Tool was created for leagueoflegends-git setup, but it *may* work with Lutris too.
This tool simply looks for working processes every 5 (value `check_delay` in config) seconds and checks for League of Legends. When tool finds game it sets "gaps" between displays, by default 100px (value `offset` in settings). With this gaps you must make very fast movement to move cursor between displays (bigger gaps = needed faster move) - then you can't move cursor outside game by accident. When tool detect that game was closed it reverts monitor's layout fron configuration file.

## Usage
1. Download latest release from Github and put into some directory (etc. /opt/league-of-hyprland)
2. Create file `config.json` in the same directory
3. Copy values from example config to your file
4. Put correct values for your setup - same as you did in Hyprland. Set `game_display` to `true` on display for playing League
5. Add tool as `exec-once` in your Hyprland configuration (working directory *MUST* be directory of tool)

## Stability
I tested tool with my setups (3 displays, everyone 1440p) with game in center - works great for me. If you experience any problems please create an Issue or even better - pull request with fix.
