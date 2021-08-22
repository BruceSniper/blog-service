create table blog_article_tag (
    `id` int(10) unsigned not null auto_increment,
    `article_id` int(11) not null comment '文章 ID',
    `tag_id` int(10) unsigned not null default 0 comment '标签 ID',
    `created_on` int unsigned default 0 not null comment '创建时间',
	`create_by` varchar(100) default '' not null comment '创建人',
	`modified_on` int unsigned default 0 not null comment '修改时间',
	`modified_by` varchar(100) default '' not null comment '修改人',
	`deleted_on` int unsigned default 0 not null comment '删除时间',
	`is_del` tinyint unsigned default 0 not null comment '是否删除 0 为未删除，1 为已删除',
    primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='文章标签关联';