# Write your MySQL query statement below
Select 
    m.name as name
From
    Employee e
Left Join
    Employee m
On
    e.managerId = m.id
Where
    m.id IS NOT NULL
Group By
    m.id
Having
    Count(e.managerId) >= 5