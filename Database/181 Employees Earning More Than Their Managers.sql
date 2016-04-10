# 这是一张员工表,表中有属性id,姓名,薪水,以及他们领导的id
# 要求在表中找出比他们领导的薪水还高的员工的姓名

# 用自连接可以解决该题目

# Write your MySQL query statement below
SELECT a.name FROM Employee a,Employee b
WHERE a.ManagerId = b.Id AND a.Salary > b.Salary;