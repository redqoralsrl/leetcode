-- 내가 푼거 속도 느림
select id from weather w
where w.temperature > (
    select s.temperature from weather s
    where s.recordDate = date_trunc('day', w.recordDate) - interval '1 day'
);
-- 테이블 조인 시 속도 상승
select w.id from weather w
inner join weather w2 on w2.recordDate = date_trunc('day', w.recordDate) - interval '1 day'
where w.temperature > w2.temperature