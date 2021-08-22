create table `xxx` (
	`created_on` int unsigned default 0 not null comment '创建时间',
	`create_by` varchar(100) default '' not null comment '创建人',
	`modified_on` int unsigned default 0 not null comment '修改时间',
	`modified_by` varchar(100) default '' not null comment '修改人',
	`deleted_on` int unsigned default 0 not null comment '删除时间',
	`is_del` tinyint unsigned default 0 not null comment '是否删除 0 为未删除，1 为已删除'
)