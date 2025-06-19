# Write your MySQL query statement below
Select
    product_id,
    Ifnull(Round(Sum(sales) / Sum(units), 2), 0) AS average_price
From
    (
        Select
            Prices.product_id AS product_id,
            Prices.price * UnitsSold.units AS sales,
            UnitsSold.units AS units
        From
            Prices
            Left Join UnitsSold On Prices.product_id = UnitsSold.product_id
            And (
                UnitsSold.purchase_date Between Prices.start_date
                And Prices.end_date
            )
    ) T
Group By
    product_id