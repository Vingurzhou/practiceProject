CREATE TABLE IF NOT EXISTS Employee (Id INT, Salary INT);
TRUNCATE TABLE Employee;
INSERT INTO Employee (id, salary) VALUES (1, 100);
INSERT INTO Employee (id, salary) VALUES (2, 200);
INSERT INTO Employee (id, salary) VALUES (3, 300);

DROP FUNCTION IF EXISTS getNthHighestSalary;
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
    READS SQL DATA
BEGIN
    set N=N-1;
    RETURN (
        # Write your MySQL query statement below.
select ifnull((select distinct salary from Employee order by Salary desc limit 1 offset N),null)
        );
END ;

SELECT getNthHighestSalary(2) AS nthHighestSalary;
