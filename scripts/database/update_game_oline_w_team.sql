UPDATE ml.game_oline_w_team
SET    
team = A.team
FROM   ml.career_oline
JOIN   B ON A.id = B.id
WHERE  C.id = A.id 