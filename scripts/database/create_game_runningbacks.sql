CREATE SCHEMA IF NOT EXISTS ml; 

CREATE TABLE IF NOT EXISTS ml.game_runningbacks(); 

ALTER TABLE ml.game_runningbacks ADD player_id TEXT;
ALTER TABLE ml.game_runningbacks ADD player_name TEXT;
ALTER TABLE ml.game_runningbacks ADD position TEXT;
ALTER TABLE ml.game_runningbacks ADD nfl_year TEXT;
ALTER TABLE ml.game_runningbacks ADD season TEXT;
ALTER TABLE ml.game_runningbacks ADD week TEXT;
ALTER TABLE ml.game_runningbacks ADD game_date TEXT;
ALTER TABLE ml.game_runningbacks ADD home_or_away TEXT;
ALTER TABLE ml.game_runningbacks ADD opponent TEXT;
ALTER TABLE ml.game_runningbacks ADD outcome TEXT;
ALTER TABLE ml.game_runningbacks ADD score TEXT;
ALTER TABLE ml.game_runningbacks ADD games_played TEXT;
ALTER TABLE ml.game_runningbacks ADD games_started TEXT;
ALTER TABLE ml.game_runningbacks ADD receptions TEXT;
ALTER TABLE ml.game_runningbacks ADD receiving_yards TEXT;
ALTER TABLE ml.game_runningbacks ADD yards_per_reception TEXT;
ALTER TABLE ml.game_runningbacks ADD longest_reception TEXT;
ALTER TABLE ml.game_runningbacks ADD receiving_tds TEXT;
ALTER TABLE ml.game_runningbacks ADD rushing_attempts TEXT;
ALTER TABLE ml.game_runningbacks ADD rushing_yards TEXT;
ALTER TABLE ml.game_runningbacks ADD yards_per_carry TEXT;
ALTER TABLE ml.game_runningbacks ADD longest_rushing_run TEXT;
ALTER TABLE ml.game_runningbacks ADD rushing_tds TEXT;
ALTER TABLE ml.game_runningbacks ADD fumbles TEXT;
ALTER TABLE ml.game_runningbacks ADD fumbles_lost TEXT;

COPY ml.game_runningbacks
(player_id,player_name,position,nfl_year,season,week,game_date,home_or_away,opponent,outcome,score,games_played,games_started,receptions,receiving_yards,yards_per_reception,longest_reception,receiving_tds,rushing_attempts,rushing_yards,yards_per_carry,longest_rushing_run,rushing_tds,fumbles,fumbles_lost) FROM '/workspace/data/Game_Logs_Runningback.csv' CSV HEADER DELIMITER ',';

UPDATE ml.game_runningbacks SET position = 'RB' WHERE true;