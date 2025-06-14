# Write your MySQL query statement below
Select * From
    cinema
Where
    description != 'boring' And
    id % 2 = 1 
Order By 
    rating Desc