Create table If Not Exists Logs
(
    id  int,
    num int
);
Truncate table Logs;
insert into Logs (id, num)
values ('1', '1');
insert into Logs (id, num)
values ('2', '1');
insert into Logs (id, num)
values ('3', '1');
insert into Logs (id, num)
values ('4', '2');
insert into Logs (id, num)
values ('5', '1');
insert into Logs (id, num)
values ('6', '2');
insert into Logs (id, num)
values ('7', '2');
# Write your MySQL query statement below
SELECT distinct l1.num as ConsecutiveNums FROM Logs l1 INNER JOIN Logs l2 ON l1.Id = l2.Id - 1 AND l1.Num = l2.Num INNER JOIN Logs l3 ON l2.Id = l3.Id - 1 AND l2.Num = l3.Num;


