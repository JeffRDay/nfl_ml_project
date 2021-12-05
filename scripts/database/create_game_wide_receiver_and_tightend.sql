CREATE SCHEMA IF NOT EXISTS ml; 

CREATE TABLE IF NOT EXISTS ml.game_wr_and_te(); 

ALTER TABLE ml.game_wr_and_te ADD player_id TEXT;
ALTER TABLE ml.game_wr_and_te ADD player_name TEXT;
ALTER TABLE ml.game_wr_and_te ADD position TEXT;
ALTER TABLE ml.game_wr_and_te ADD nfl_year TEXT;
ALTER TABLE ml.game_wr_and_te ADD season TEXT;
ALTER TABLE ml.game_wr_and_te ADD week TEXT;
ALTER TABLE ml.game_wr_and_te ADD game_date TEXT;
ALTER TABLE ml.game_wr_and_te ADD home_or_away TEXT;
ALTER TABLE ml.game_wr_and_te ADD opponent TEXT;
ALTER TABLE ml.game_wr_and_te ADD outcome TEXT;
ALTER TABLE ml.game_wr_and_te ADD score TEXT;
ALTER TABLE ml.game_wr_and_te ADD games_played TEXT;
ALTER TABLE ml.game_wr_and_te ADD games_started TEXT;
ALTER TABLE ml.game_wr_and_te ADD receptions TEXT;
ALTER TABLE ml.game_wr_and_te ADD receiving_yards TEXT;
ALTER TABLE ml.game_wr_and_te ADD yards_per_reception TEXT;
ALTER TABLE ml.game_wr_and_te ADD longest_reception TEXT;
ALTER TABLE ml.game_wr_and_te ADD receiving_tds TEXT;
ALTER TABLE ml.game_wr_and_te ADD rushing_attempts TEXT;
ALTER TABLE ml.game_wr_and_te ADD rushing_yards TEXT;
ALTER TABLE ml.game_wr_and_te ADD yards_per_carry TEXT;
ALTER TABLE ml.game_wr_and_te ADD longest_rushing_run TEXT;
ALTER TABLE ml.game_wr_and_te ADD rushing_tds TEXT;
ALTER TABLE ml.game_wr_and_te ADD fumbles TEXT;
ALTER TABLE ml.game_wr_and_te ADD fumbles_lost TEXT;

COPY ml.game_wr_and_te
(player_id,player_name,position,nfl_year,season,week,game_date,home_or_away,opponent,outcome,score,games_played,games_started,receptions,receiving_yards,yards_per_reception,longest_reception,receiving_tds,rushing_attempts,rushing_yards,yards_per_carry,longest_rushing_run,rushing_tds,fumbles,fumbles_lost) FROM '/workspace/data/Game_Logs_Wide_Receiver_and_Tight_End.csv' CSV HEADER DELIMITER ',';