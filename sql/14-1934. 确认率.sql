Select 
    s.user_id,
    ROUND(IFNULL(AVG(c.action = 'confirmed'), 0), 2) AS confirmation_rate
From
    Signups s
Left Join
    Confirmations c
On
    s.user_id = c.user_id
Group By
    s.user_id;