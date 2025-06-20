# Write your MySQL query statement below
Select
    project_id,
    Round(AVG(experience_years), 2) As average_years
From
    Project
    Left Join Employee On Project.employee_id = Employee.employee_id
Group By
    project_id