CREATE TABLE `douban_books` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(30) DEFAULT '' COMMENT '标题',
  `author` varchar(50) DEFAULT '' COMMENT '作者',
  `translator` varchar(50) DEFAULT '' COMMENT '译者',
  `press` varchar(50) DEFAULT '' COMMENT '出版社',
  `date` varchar(20) DEFAULT '' COMMENT '日期',
  `price` float DEFAULT '0.0' COMMENT '定价',
  `star` int(10) unsigned DEFAULT '0' COMMENT 'star',
  `comment` int(10) unsigned DEFAULT '0' COMMENT '评分',
  `quote` varchar(999) DEFAULT '' COMMENT '引用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='豆瓣读书Top250';
