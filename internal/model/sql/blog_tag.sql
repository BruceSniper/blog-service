create table `blog_tag` (
    `id` int(10) unsigned not null auto_increment,
	`name` varchar(100) default '' comment '标签名称',
	`created_on` int unsigned default 0 not null comment '创建时间',
	`create_by` varchar(100) default '' not null comment '创建人',
	`modified_on` int unsigned default 0 not null comment '修改时间',
	`modified_by` varchar(100) default '' not null comment '修改人',
	`deleted_on` int unsigned default 0 not null comment '删除时间',
	`is_del` tinyint unsigned default 0 not null comment '是否删除 0 为未删除，1 为已删除',
	`state` tinyint(3) unsigned default 1 comment '状态 0 为禁用， 1 为启用',
	primary key (`id`)
) engine=InnoDB DEFAULT CHARSET=utf8mb4 comment='标签管理';