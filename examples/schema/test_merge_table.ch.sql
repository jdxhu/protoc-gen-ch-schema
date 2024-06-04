CREATE TABLE IF NOT EXISTS test_merge_table
(
	a Int32 COMMENT 'Description of field a -- this is an int32',
	b_a Array(Int32),
	b_b Array(String),
	c Array(String) COMMENT 'Repeated c string',
	e DateTime COMMENT 'TIMESTAMP (uint64 in proto) - required in ClickHouse',
	wkt1 Int32,
	wkt2 DateTime,
	ctime DateTime DEFAULT now() COMMENT 'create time'
) ENGINE = MergeTree()
ORDER BY ctime
PARTITION BY toYYYYMMDD(ctime)
TTL ctime + INTERVAL 3 MONTH
SETTINGS index_granularity = '8192';