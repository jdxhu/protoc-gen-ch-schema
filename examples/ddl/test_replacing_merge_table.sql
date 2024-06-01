CREATE TABLE IF NOT EXISTS test_replacing_merge_table
(
	a Int32 COMMENT 'Description of field a -- this is an int32',
	b_a Array(Int32),
	b_b Array(String),
	c Array(String) COMMENT 'Repeated c string',
	e DateTime COMMENT 'TIMESTAMP (uint64 in proto) - required in ClickHouse',
	wkt1 Int32,
	wkt2 DateTime,
	ctime DateTime DEFAULT now() COMMENT 'create time',
	deleted UInt8 COMMENT 'state: 0, deleted: 1'
) ENGINE = ReplacingMergeTree(e, deleted)
ORDER BY ctime
PARTITION BY toYYYYMM(ctime)
TTL ctime + INTERVAL 3 MONTH
SETTINGS index_granularity = '8192';