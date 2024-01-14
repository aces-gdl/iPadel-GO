select
	tg.ID as game_id,
	tg.tournament_id ,
	tgrp.group_number ,
	tgrp.id as tournament_group_id,
	c.id as category_id,
	c.color as category_color,
	c.description as category_description,  
	tt1.member1_id as team1_member1_id,
	tt1.name1 as team1_name1,
	tt1.ranking1 as team1_ranking1,
	tt1.member2_id as team1_member2_id,
	tt1.name2 as team1_name2,
	tt1.ranking2 as team1_ranking2,
	tt2.member1_id as team2_member1_id,
	tt2.name1 as team2_name1,
	tt2.ranking1 as team2_ranking1,
	tt2.member2_id as team2_member2_id,
	tt2.name2 as team2_name2,
	tt2.ranking2 as team2_ranking2,
	tg.tournament_time_slots_id,
	gr.id,
	gr.team1_set1,
	gr.team1_set2,
	gr.team1_set3, 
	gr.team2_set1,
	gr.team2_set2,
	gr.team2_set3 
from tournament_games tg 
	inner join tournament_teams tt1 on tg.team1_id = tt1.id and tg.tournament_id = tt1.tournament_id 
	inner join tournament_teams tt2 on tg.team2_id = tt2.id and tg.tournament_id = tt2.tournament_id 
	inner join categories c on tg.category_id = c.id 
	inner join tournament_groups tgrp on tg.tournament_group_id = tgrp.id
	left outer join tournament_game_results gr on tg.id = gr.game_id
where tg.tournament_id ='8196138d-9c50-4143-b786-7c3be6815435' 
						
						
						
update tournament_games set tournament_time_slots_id ='00000000-0000-0000-0000-000000000000';
update tournament_time_slots set game_id ='00000000-0000-0000-0000-000000000000';
	
							
select count(*) from tournament_games tg group by tournament_group_id ;

delete from public.tournament_games;
delete from tournament_groups ;
delete from tournament_team_by_groups;