CREATE SCHEMA IF NOT EXISTS ml; 

CREATE TABLE IF NOT EXISTS ml.game_oline_w_team(); 

ALTER TABLE ml.game_oline_w_team ADD player_id TEXT ;
ALTER TABLE ml.game_oline_w_team ADD player_name TEXT ;
ALTER TABLE ml.game_oline_w_team ADD team TEXT ;
ALTER TABLE ml.game_oline_w_team ADD position TEXT ;
ALTER TABLE ml.game_oline_w_team ADD nfl_year TEXT ;
ALTER TABLE ml.game_oline_w_team ADD season TEXT ;
ALTER TABLE ml.game_oline_w_team ADD week TEXT ;
ALTER TABLE ml.game_oline_w_team ADD game_date TEXT ;
ALTER TABLE ml.game_oline_w_team ADD home_or_away TEXT ;
ALTER TABLE ml.game_oline_w_team ADD opponent TEXT ;
ALTER TABLE ml.game_oline_w_team ADD outcome TEXT ;
ALTER TABLE ml.game_oline_w_team ADD score TEXT ;
ALTER TABLE ml.game_oline_w_team ADD games_played TEXT ;
ALTER TABLE ml.game_oline_w_team ADD games_started TEXT ;

COPY ml.game_oline_w_team
(player_id,player_name,position,nfl_year,season,week,game_date,home_or_away,opponent,outcome,score,games_played,games_started) 
FROM '/workspace/data/Game_logs_Offensive_Line.csv' CSV HEADER DELIMITER ',';