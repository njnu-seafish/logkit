package reader

import (
	"strings"

	"github.com/qiniu/logkit/utils"
)

// ModeUsages 用途说明
var ModeUsages = []utils.KeyValue{
	{ModeDir, "从文件读取( dir 模式)"},
	{ModeFile, "从文件读取( file 模式)"},
	{ModeTailx, "从文件读取( tailx 模式)"},
	{ModeMysql, "从 MySQL 读取"},
	{ModeMssql, "从 MSSQL 读取"},
	{ModeElastic, "从 Elasticsearch 读取"},
	{ModeMongo, "从 MongoDB 读取"},
	{ModeKafka, "从 kafka 读取"},
	{ModeRedis, "从 redis 读取"},
}

var (
	OptionDataSourceTag = utils.Option{
		KeyName:      KeyDataSourceTag,
		ChooseOnly:   false,
		Default:      "",
		DefaultNoUse: false,
		Description:  "具体的数据文件路径来源标签(datasource_tag)",
	}
)
var ModeKeyOptions = map[string][]utils.Option{
	ModeDir: {
		{
			KeyName:      KeyLogPath,
			ChooseOnly:   false,
			Default:      "/home/users/john/log/",
			DefaultNoUse: true,
			Description:  "日志文件夹路径(log_path)",
		},
		{
			KeyName:      KeyMetaPath,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "断点续传元数据路径(meta_path)",
		},
		{
			KeyName:      KeyFileDone,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "读取过的文件信息保存路径(file_done)",
		},
		{
			KeyName:      KeyBufSize,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "文件缓存数据大小(reader_buf_size)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:       KeyWhence,
			ChooseOnly:    true,
			ChooseOptions: []string{WhenceOldest, WhenceNewest},
			Description:   "读取的起始位置(read_from)",
		},
		{
			KeyName:    KeyEncoding,
			ChooseOnly: true,
			ChooseOptions: []string{"UTF-8", "UTF-16", "US-ASCII", "ISO-8859-1",
				"GBK", "GB18030", "EUC-JP", "UTF-16BE", "UTF-16LE", "Big5", "Shift_JIS",
				"ISO-8859-2", "ISO-8859-3", "ISO-8859-4", "ISO-8859-5", "ISO-8859-6", "ISO-8859-7",
				"ISO-8859-8", "ISO-8859-9", "ISO-8859-10", "ISO-8859-11", "ISO-8859-12", "ISO-8859-13",
				"ISO-8859-14", "ISO-8859-15", "ISO-8859-16", "macos-0_2-10.2", "macos-6_2-10.4",
				"macos-7_3-10.2", "macos-29-10.2", "macos-35-10.2", "windows-1250", "windows-1251",
				"windows-1252", "windows-1253", "windows-1254", "windows-1255", "windows-1256",
				"windows-1257", "windows-1258", "windows-874", "IBM037", "ibm-273_P100-1995",
				"ibm-277_P100-1995", "ibm-278_P100-1995", "ibm-280_P100-1995", "ibm-284_P100-1995",
				"ibm-285_P100-1995", "ibm-290_P100-1995", "ibm-297_P100-1995", "ibm-420_X120-1999",
				//此处省略大量IBM的字符集，太多，等用户需要再加
				"KOI8-R", "KOI8-U", "ebcdic-xml-us"},
			Default:      "UTF-8",
			DefaultNoUse: false,
			Description:  "编码方式(encoding)",
		},
		OptionDataSourceTag,
		{
			KeyName:      KeyReadIOLimit,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "读取速度限制(MB/s)(readio_limit)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:      KeyHeadPattern,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "多行读取的起始行正则表达式(head_pattern)",
		},
		{
			KeyName:       KeyIgnoreHiddenFile,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "是否忽略隐藏文件(ignore_hidden)",
		},
		{
			KeyName:      KeyIgnoreFileSuffix,
			ChooseOnly:   false,
			Default:      strings.Join(defaultIgnoreFileSuffix, ","),
			DefaultNoUse: false,
			Description:  "根据后缀忽略文件(ignore_file_suffix)",
		},
		{
			KeyName:      KeyValidFilePattern,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "根据正则表达式匹配文件(valid_file_pattern)",
		},
	},
	ModeFile: {
		{
			KeyName:      KeyLogPath,
			ChooseOnly:   false,
			Default:      "/home/users/john/log/my.log",
			DefaultNoUse: true,
			Description:  "日志文件路径(log_path)",
		},
		{
			KeyName:      KeyMetaPath,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "断点续传元数据路径(meta_path)",
		},
		{
			KeyName:      KeyBufSize,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "文件缓存数据大小(reader_buf_size)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:       KeyWhence,
			ChooseOnly:    true,
			ChooseOptions: []string{WhenceOldest, WhenceNewest},
			Description:   "读取的起始位置(read_from)",
		},
		OptionDataSourceTag,
		{
			KeyName:    KeyEncoding,
			ChooseOnly: true,
			ChooseOptions: []string{"UTF-8", "UTF-16", "US-ASCII", "ISO-8859-1",
				"GBK", "GB18030", "EUC-JP", "UTF-16BE", "UTF-16LE", "Big5", "Shift_JIS",
				"ISO-8859-2", "ISO-8859-3", "ISO-8859-4", "ISO-8859-5", "ISO-8859-6", "ISO-8859-7",
				"ISO-8859-8", "ISO-8859-9", "ISO-8859-10", "ISO-8859-11", "ISO-8859-12", "ISO-8859-13",
				"ISO-8859-14", "ISO-8859-15", "ISO-8859-16", "macos-0_2-10.2", "macos-6_2-10.4",
				"macos-7_3-10.2", "macos-29-10.2", "macos-35-10.2", "windows-1250", "windows-1251",
				"windows-1252", "windows-1253", "windows-1254", "windows-1255", "windows-1256",
				"windows-1257", "windows-1258", "windows-874", "IBM037", "ibm-273_P100-1995",
				"ibm-277_P100-1995", "ibm-278_P100-1995", "ibm-280_P100-1995", "ibm-284_P100-1995",
				"ibm-285_P100-1995", "ibm-290_P100-1995", "ibm-297_P100-1995", "ibm-420_X120-1999",
				//省略大量IBM的字符集，太多，等用户需要再加
				"KOI8-R", "KOI8-U", "ebcdic-xml-us"},
			Default:      "UTF-8",
			DefaultNoUse: false,
			Description:  "编码方式(encoding)",
		},
		{
			KeyName:      KeyReadIOLimit,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "读取速度限制(MB/s)(readio_limit)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:      KeyHeadPattern,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "多行读取的起始行正则表达式(head_pattern)",
		},
	},
	ModeTailx: {
		{
			KeyName:      KeyLogPath,
			ChooseOnly:   false,
			Default:      "/home/users/*/mylog/*.log",
			DefaultNoUse: true,
			Description:  "日志文件路径模式串(log_path)",
		},
		{
			KeyName:      KeyMetaPath,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "断点续传元数据路径(meta_path)",
		},
		{
			KeyName:      KeyBufSize,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "文件缓存数据大小(reader_buf_size)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:       KeyWhence,
			ChooseOnly:    true,
			ChooseOptions: []string{WhenceOldest, WhenceNewest},
			Description:   "读取的起始位置(read_from)",
		},
		{
			KeyName:    KeyEncoding,
			ChooseOnly: true,
			ChooseOptions: []string{"UTF-8", "UTF-16", "US-ASCII", "ISO-8859-1",
				"GBK", "GB18030", "EUC-JP", "UTF-16BE", "UTF-16LE", "Big5", "Shift_JIS",
				"ISO-8859-2", "ISO-8859-3", "ISO-8859-4", "ISO-8859-5", "ISO-8859-6", "ISO-8859-7",
				"ISO-8859-8", "ISO-8859-9", "ISO-8859-10", "ISO-8859-11", "ISO-8859-12", "ISO-8859-13",
				"ISO-8859-14", "ISO-8859-15", "ISO-8859-16", "macos-0_2-10.2", "macos-6_2-10.4",
				"macos-7_3-10.2", "macos-29-10.2", "macos-35-10.2", "windows-1250", "windows-1251",
				"windows-1252", "windows-1253", "windows-1254", "windows-1255", "windows-1256",
				"windows-1257", "windows-1258", "windows-874", "IBM037", "ibm-273_P100-1995",
				"ibm-277_P100-1995", "ibm-278_P100-1995", "ibm-280_P100-1995", "ibm-284_P100-1995",
				"ibm-285_P100-1995", "ibm-290_P100-1995", "ibm-297_P100-1995", "ibm-420_X120-1999",
				//省略大量IBM的字符集，太多，等用户需要再加
				"KOI8-R", "KOI8-U", "ebcdic-xml-us"},
			Default:      "UTF-8",
			DefaultNoUse: false,
			Description:  "编码方式(encoding)",
		},
		{
			KeyName:      KeyReadIOLimit,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "读取速度限制(MB/s)(readio_limit)",
			CheckRegex:   "\\d+",
		},
		OptionDataSourceTag,
		{
			KeyName:      KeyHeadPattern,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "多行读取的起始行正则表达式(head_pattern)",
		},
		{
			KeyName:      KeyExpire,
			ChooseOnly:   false,
			Default:      "24h",
			DefaultNoUse: false,
			Description:  "文件过期时间(时h，分m，秒s)(expire)",
			CheckRegex:   "\\d+[hms]",
		},
		{
			KeyName:      KeyMaxOpenFiles,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "最大的打开文件数(max_open_files)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:      KeyStatInterval,
			ChooseOnly:   false,
			Default:      "3m",
			DefaultNoUse: false,
			Description:  "文件扫描间隔(stat_interval)",
			CheckRegex:   "\\d+[hms]",
		},
	},
	ModeMysql: {
		{
			KeyName:      KeyMysqlDataSource,
			ChooseOnly:   false,
			Default:      "<username>:<password>@tcp(<hostname>:<port>)",
			DefaultNoUse: true,
			Description:  "数据库地址(mysql_datasource)",
		},
		{
			KeyName:      KeyMysqlDataBase,
			ChooseOnly:   false,
			Default:      "<database>",
			DefaultNoUse: true,
			Description:  "数据库名称(mysql_database)",
		},
		{
			KeyName:      KeyMysqlSQL,
			ChooseOnly:   false,
			Default:      "select * from <table>;",
			DefaultNoUse: true,
			Description:  "数据查询语句(mysql_sql)",
		},
		{
			KeyName:      KeyMysqlOffsetKey,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "递增的列名称(mysql_offset_key)",
		},
		{
			KeyName:      KeyMysqlReadBatch,
			ChooseOnly:   false,
			Default:      "100",
			DefaultNoUse: false,
			Description:  "分批查询的单批次大小(mysql_limit_batch)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:      KeyMetaPath,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "断点续传元数据路径(meta_path)",
		},
		OptionDataSourceTag,
		{
			KeyName:      KeyMysqlCron,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "定时任务调度Cron(mysql_cron)",
		},
		{
			KeyName:       KeyMysqlExecOnStart,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "启动时立即执行(mysql_exec_onstart)",
		},
		{
			KeyName:      KeySQLSchema,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "SQL字段类型定义(sql_schema)",
		},
	},
	ModeMssql: {
		{
			KeyName:      KeyMssqlDataSource,
			ChooseOnly:   false,
			Default:      "server=<hostname or instance>;user id=<username>;password=<password>;port=<port>",
			DefaultNoUse: true,
			Description:  "数据库地址(mssql_datasource)",
		},
		{
			KeyName:      KeyMssqlDataBase,
			ChooseOnly:   false,
			Default:      "<database>",
			DefaultNoUse: true,
			Description:  "数据库名称(mssql_database)",
		},
		{
			KeyName:      KeyMssqlSQL,
			ChooseOnly:   false,
			Default:      "select * from <table>;",
			DefaultNoUse: true,
			Description:  "数据查询语句(mssql_sql)",
		},
		{
			KeyName:      KeyMssqlOffsetKey,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: true,
			Description:  "递增的列名称(mssql_offset_key)",
		},
		{
			KeyName:      KeyMetaPath,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "断点续传元数据路径(meta_path)",
		},
		OptionDataSourceTag,
		{
			KeyName:      KeyMssqlReadBatch,
			ChooseOnly:   false,
			Default:      "100",
			DefaultNoUse: false,
			Description:  "分批查询的单批次大小(mssql_limit_batch)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:      KeyMssqlCron,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "定时任务调度Crontab(mssql_cron)",
		},
		{
			KeyName:       KeyMssqlExecOnStart,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "启动时立即执行(mssql_exec_onstart)",
		},
		{
			KeyName:      KeySQLSchema,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "SQL字段类型定义(sql_schema)",
		},
	},
	ModeElastic: {
		{
			KeyName:      KeyESHost,
			ChooseOnly:   false,
			Default:      "http://localhost:9200",
			DefaultNoUse: true,
			Description:  "数据库地址(es_host)",
		},
		{
			KeyName:       KeyESVersion,
			ChooseOnly:    true,
			ChooseOptions: []string{ElasticVersion2, ElasticVersion5},
			Description:   "ES版本号(es_version)",
		},
		{
			KeyName:      KeyESIndex,
			ChooseOnly:   false,
			Default:      "app-repo-123",
			DefaultNoUse: true,
			Description:  "ES索引名称(es_index)",
		},
		{
			KeyName:      KeyESType,
			ChooseOnly:   false,
			Default:      "type_app",
			DefaultNoUse: true,
			Description:  "ES的app名称(es_type)",
		},
		{
			KeyName:      KeyMetaPath,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "断点续传元数据路径(meta_path)",
		},
		OptionDataSourceTag,
		{
			KeyName:      KeyESReadBatch,
			ChooseOnly:   false,
			Default:      "100",
			DefaultNoUse: false,
			Description:  "分批查询的单批次大小(es_limit_batch)",
		},
		{
			KeyName:      KeyESKeepAlive,
			ChooseOnly:   false,
			Default:      "1d",
			DefaultNoUse: false,
			Description:  "ES的Offset保存时间(es_keepalive)",
			CheckRegex:   "\\d+[dms]",
		},
	},
	ModeMongo: {
		{
			KeyName:      KeyMongoHost,
			ChooseOnly:   false,
			Default:      "mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]",
			DefaultNoUse: true,
			Description:  "数据库地址(mongo_host)",
		},
		{
			KeyName:      KeyMongoDatabase,
			ChooseOnly:   false,
			Default:      "app123",
			DefaultNoUse: true,
			Description:  "数据库名称(mongo_database)",
		},
		{
			KeyName:      KeyMongoCollection,
			ChooseOnly:   false,
			Default:      "collection1",
			DefaultNoUse: true,
			Description:  "数据表名称(mongo_collection)",
		},
		{
			KeyName:      KeyMongoOffsetKey,
			ChooseOnly:   false,
			Default:      "_id",
			DefaultNoUse: true,
			Description:  "递增的主键(mongo_offset_key)",
		},
		{
			KeyName:      KeyMetaPath,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "断点续传元数据路径(meta_path)",
		},
		OptionDataSourceTag,
		{
			KeyName:      KeyMongoReadBatch,
			ChooseOnly:   false,
			Default:      "100",
			DefaultNoUse: false,
			Description:  "分批查询的单批次大小(mongo_limit_batch)",
			CheckRegex:   "\\d+",
		},
		{
			KeyName:      KeyMongoCron,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "定时任务调度Cron(mongo_cron)",
		},
		{
			KeyName:       KeyMongoExecOnstart,
			ChooseOnly:    true,
			ChooseOptions: []string{"true", "false"},
			Default:       "true",
			DefaultNoUse:  false,
			Description:   "启动时立即执行(mongo_exec_onstart)",
		},
		{
			KeyName:      KeyMongoFilters,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "数据过滤方式(mongo_filters)",
		},
	},
	ModeKafka: {
		{
			KeyName:      KeyKafkaGroupID,
			ChooseOnly:   false,
			Default:      "logkit1",
			DefaultNoUse: true,
			Description:  "Kafka的consumer组名称(kafka_groupid)",
		},
		{
			KeyName:      KeyKafkaTopic,
			ChooseOnly:   false,
			Default:      "test_topic1",
			DefaultNoUse: true,
			Description:  "Kafka的topic名称(kafka_topic)",
		},
		{
			KeyName:      KeyKafkaZookeeper,
			ChooseOnly:   false,
			Default:      "localhost:2181",
			DefaultNoUse: true,
			Description:  "Zookeeper地址(kafka_zookeeper)",
		},
		{
			KeyName:       KeyWhence,
			ChooseOnly:    true,
			ChooseOptions: []string{WhenceOldest, WhenceNewest},
			Description:   "读取的起始位置(read_from)",
		},
		OptionDataSourceTag,
	},
	ModeRedis: {
		{
			KeyName:       KeyRedisDataType,
			ChooseOnly:    true,
			ChooseOptions: []string{DataTypeList, DataTypeChannel, DataTypePatterChannel},
			Description:   "Redis的数据读取模式(redis_datatype)",
		},
		{
			KeyName:      KeyRedisDB,
			ChooseOnly:   false,
			Default:      "0",
			DefaultNoUse: true,
			Description:  "数据库名称(redis_db)",
		},
		{
			KeyName:      KeyRedisKey,
			ChooseOnly:   false,
			Default:      "key1",
			DefaultNoUse: true,
			Description:  "redis键(redis_key)",
		},
		{
			KeyName:      KeyRedisAddress,
			ChooseOnly:   false,
			Default:      "127.0.0.1:6379",
			DefaultNoUse: false,
			Description:  "数据库地址(redis_address)",
		},
		{
			KeyName:      KeyRedisPassword,
			ChooseOnly:   false,
			Default:      "",
			DefaultNoUse: false,
			Description:  "密码(redis_password)",
		},
		{
			KeyName:      KeyTimeoutDuration,
			ChooseOnly:   false,
			Default:      "5s",
			DefaultNoUse: false,
			Description:  "单次读取超时时间(m(分)、s(秒))(redis_timeout)",
			CheckRegex:   "\\d+[ms]",
		},
		OptionDataSourceTag,
	},
}
