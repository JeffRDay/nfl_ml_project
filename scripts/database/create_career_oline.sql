CREATE SCHEMA IF NOT EXISTS ml; 

CREATE TABLE IF NOT EXISTS ml.career_oline();

ALTER TABLE ml.career_oline ADD player_id TEXT ;
ALTER TABLE ml.career_oline ADD player_name TEXT ;
ALTER TABLE ml.career_oline ADD position TEXT ;
ALTER TABLE ml.career_oline ADD nfl_year TEXT ;
ALTER TABLE ml.career_oline ADD team TEXT ;
ALTER TABLE ml.career_oline ADD games_played TEXT ;
ALTER TABLE ml.career_oline ADD games_started TEXT ;

COPY ml.career_oline
(player_id,player_name,position,nfl_year,team,games_played,games_started)
FROM '/workspace/data/Career_Stats_Offensive_Line.csv'
CSV HEADER DELIMITER ',';