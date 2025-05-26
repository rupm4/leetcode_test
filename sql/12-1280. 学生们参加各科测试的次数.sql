# Write your MySQL query statement below
Select
    s.student_id,
    s.student_name,
    sub.subject_name,
    count(e.student_id) as attended_exams
From
    Students s
Join
    Subjects sub
Left Join 
    Examinations e
On
    e.student_id = s.student_id and e.subject_name = sub.subject_name
Group By
    s.student_id, sub.subject_name
Order By
    s.student_id, sub.subject_name
