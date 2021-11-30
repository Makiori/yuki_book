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
	`password` varchar(255) not null comment '用户密码',
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

drop table if exists `user_type`;
create table `user_type` (
	`id` int not null AUTO_INCREMENT comment '类型id',
	`type_name` varchar(255) not null comment '用户类型',
	`Max_bor_num` int not null comment '借书上限',
	`Max_time`  int not null comment '借书最长时间',
	`Max_count` int not null comment '最大续借时间',
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  	primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='用户类型';


drop table if exists `book`;
create table `book` (
	`id` int not null  comment '书本id',
	`book_class_id` int not null comment '书籍id',
	`shelfNum` int not null comment '书架编号id',
	`bookState` varchar(255) comment '书本状态',
	`bookDamage` varchar(255) comment '书本损害程度',
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  	primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='书本';

drop table if exists `book_class`;
create table `book_class` (
	`id` int not null  comment '书集id',
	`bookName` varchar(255) not null comment '书集名字'
	`author` varchar(255) not null comment '书集作者',
	`bookKey` varchar(255) not null comment '书集主题',
	`bookEdit` varchar(255) not null comment '书集出版社',
	`pageNum` int not null comment '书集页数',
	`bookType` int not null comment '书集类型id'
	`publishTime` datetime not null comment '出版日期',
	`bookNum` int null DEFAULT 0 comment '库存量',
	`bookIN` int null DEFAULT 0 comment '在馆数量'，
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  	primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='书集';


drop table if exists `user_borrow`;
create table `book_borrow` (
	`id` int not null  comment 'id',
	`user_id` int not null comment '用户id',
	`book_class_id` int not null comment '书集id',
	`book_id` int not null comment '书本id',
	`borrow_at` datetime not null comment '借出日期',
	`borrow_count` int not null comment '续借次数',
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  	primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='借阅记录';

drop table if exists `reading_room`;
create table `reading_room` (
	`id` int not null  comment '阅览室id',
	`name` varchar(255) not null comment '阅览室名字',
	`po` varchar(255) not null comment '阅览室位置',
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  	primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='阅览室';

drop table if exists `reading_room`;
create table `reading_room` (
	`id` int not null  comment '书架id',
	`reading_room_id` int not null comment '阅览室id',
	`name` varchar(255) not null comment '书架分类',
	`created_at` datetime not null comment '创建时间',
	`updated_at` datetime not null comment '修改时间',
  	primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='书架';