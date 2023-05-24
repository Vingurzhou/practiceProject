drop table Employee;
Create table If Not Exists Employee
(
    id        int,
    name      varchar(255),
    salary    int,
    managerId int
);
Truncate table Employee;
insert into Employee (id, name, salary, managerId)
values ('1', 'Joe', '70000', '3');
insert into Employee (id, name, salary, managerId)
values ('2', 'Henry', '80000', '4');
insert into Employee (id, name, salary, managerId)
values ('3', 'Sam', '60000', null);
insert into Employee (id, name, salary, managerId)
values ('4', 'Max', '90000', null);

# Write your MySQL query statement below
select e1.name as Employee
from Employee e1
         left join Employee e2 on e1.managerId = e2.id
where e1.salary > e2.salary