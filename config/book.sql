CREATE TABLE `douban_book` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `author` varchar(50) DEFAULT '' COMMENT '作者',
  `translator` varchar(50) DEFAULT '' COMMENT '译者',
  `press` varchar(50) DEFAULT '' COMMENT '出版社',
  `date` varchar(20) DEFAULT '' COMMENT '日期',
  `price` varchar(50) DEFAULT '' COMMENT '定价',
  `star` varchar(20) DEFAULT '' COMMENT '星级',
  `comment` int(10) unsigned DEFAULT '0' COMMENT '评分',
  `quote` varchar(999) DEFAULT '' COMMENT '引用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='豆瓣读书Top250';
