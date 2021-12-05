# nfl_ml_project
VT CS5644 Machine Learning with Big Data - Semester Project: Create a ML model for predicting fantasy football scores

## Setup

Pull pgAdmin
```bash
docker run -p 5001:80 \
    -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' \
    -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' \
    -d dpage/pgadmin4
```

Example:
player_id, Chicago Bears, Philidelphia Eagles, ..., nfl_year, preason, regular_season, post_season, home, away 
45687987        1                    0                1975       0          1                0        1     0 
 
player_id - updated to just the number
<team> - 1/0 [1 = player is on the team, else 0]
.
.
.
player_name - drop
rb - 1/0
wr - 1/0
te - 1/0
nfl_year - 1984
preseason - 1/0
regular_season - 1/0
post_season - 1/0
week - 1
game_date - drop 
home - 1/0
away - 1/0
opponent - drop (but used to identify d-line)
outcome - drop
score - drop
games_played - drop
games_started - drop
receptions
receiving_yards
yards_per_reception
longest_reception
receiving_tds
rushing_attempts
rushing_yards
yards_per_carry
longest_rushing_run
rushing_tds
fumbles
fumbles_lost
fantasy_score - 54.8
oline_<player_id> - 1/0 [1 = olineman is on team]
.
.
.
dline_<player_id> - 1/0 [1 = player played against this dline in game]
.
.
.
