drop table if exists `admin`;
create table `admin` (
	`id` int not null AUTO_INCREMENT comment '管理员id',
	`username` varchar(255) not null comment '管理员账号',
	`password` varchar(255) not null comment '管理员密码',
	`salt` varchar(255) not null comment '混淆盐', 
	`admin_name` varchar(255) comment '管理员姓名',
	`admin_phonenumber` varchar(255) not null comment '管理员电话',
	`admin_address` varchar(255) not null comment '管理员住址',
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='管理员';


drop table if exists `user`;
create table `user` (
	`id` int not null AUTO_INCREMENT comment '用户id',
	`username` varchar(255) not null comment '用户账号',
	`PassWord` varchar(255) not null comment '用户密码',
	`salt` varchar(255) not null comment '混淆盐', 
	`reader_name` varchar(255) comment '用户姓名',
	`reader_phonenumber` varchar(255) not null comment '用户电话',
	`reader_address` varchar(255) not null comment '用户住址',
	`reader_class` varchar(255) comment '用户班级',
	`reader_Email` varchar(255) not null comment '用户邮箱',
	`reader_type` varchar(255) not null comment '用户类型',
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  	primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='用户';