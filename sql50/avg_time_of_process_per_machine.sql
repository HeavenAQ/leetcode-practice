SELECT
    actStart.machine_id,
    ROUND(CAST(AVG(actEnd.timestamp - actStart.timestamp) AS NUMERIC), 3) AS processing_time
FROM
    Activity actStart
    INNER JOIN Activity actEnd ON actStart.process_id = actEnd.process_id
WHERE
    actStart.activity_type = 'start'
    AND actEnd.activity_type = 'end'
    AND actStart.machine_id = actEnd.machine_id
GROUP BY
    actStart.machine_id;
