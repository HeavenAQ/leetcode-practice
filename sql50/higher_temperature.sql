-- Write your PostgreSQL query statement below
SELECT
    w2.id AS Id
FROM
    weather w1
    INNER JOIN weather w2 ON w1.recordDate = w2.recordDate - 1
WHERE
    w1.temperature < w2.temperature;
