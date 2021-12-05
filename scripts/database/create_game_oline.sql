CREATE SCHEMA IF NOT EXISTS ml; 

CREATE TABLE IF NOT EXISTS ml.game_oline(); 

ALTER TABLE ml.game_oline ADD player_id TEXT ;
ALTER TABLE ml.game_oline ADD player_name TEXT ;
ALTER TABLE ml.game_oline ADD position TEXT ;
ALTER TABLE ml.game_oline ADD nfl_year TEXT ;
ALTER TABLE ml.game_oline ADD season TEXT ;
ALTER TABLE ml.game_oline ADD week TEXT ;
ALTER TABLE ml.game_oline ADD game_date TEXT ;
ALTER TABLE ml.game_oline ADD home_or_away TEXT ;
ALTER TABLE ml.game_oline ADD opponent TEXT ;
ALTER TABLE ml.game_oline ADD outcome TEXT ;
ALTER TABLE ml.game_oline ADD score TEXT ;
ALTER TABLE ml.game_oline ADD games_played TEXT ;
ALTER TABLE ml.game_oline ADD games_started TEXT ;

COPY ml.game_oline
(player_id,player_name,position,nfl_year,season,week,game_date,home_or_away,opponent,outcome,score,games_played,games_started) FROM '/workspace/data/Game_logs_Offensive_Line.csv' CSV HEADER DELIMITER ',';