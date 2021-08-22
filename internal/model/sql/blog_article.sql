create table `blog_article` (
    `id` int(10) unsigned not null auto_increment,
    `title` varchar(100) default '' comment '文章标题',
    `desc` varchar(255) default '' comment '文章简述',
    `cover_image_url` varchar(255) default '' comment '封面图片地址',
    `content` longtext comment '文章内容',
    `created_on` int unsigned default 0 not null comment '创建时间',
	`create_by` varchar(100) default '' not null comment '创建人',
	`modified_on` int unsigned default 0 not null comment '修改时间',
	`modified_by` varchar(100) default '' not null comment '修改人',
	`deleted_on` int unsigned default 0 not null comment '删除时间',
	`is_del` tinyint unsigned default 0 not null comment '是否删除 0 为未删除，1 为已删除',
    `state` tinyint(3) unsigned default 1 comment '状态 0 为禁用、1 为启用',
    primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='文章管理';