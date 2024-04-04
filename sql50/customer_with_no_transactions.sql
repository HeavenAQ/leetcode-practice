-- Write your PostgreSQL query statement below
SELECT
    Visits.customer_id,
    count(Visits.visit_id) AS count_no_trans
FROM
    Visits
    FULL JOIN Transactions ON Transactions.visit_id = Visits.visit_id
WHERE
    Transactions.visit_id IS NULL
GROUP BY
    customer_id;
