
-- +migrate Up
create table if not exists users (
  id bigint not null primary key auto_increment comment 'ユーザーID',
  name varchar(255) not null unique comment 'ユーザー名',
  password varchar(255) not null unique comment 'パスワード',
  role varchar(255) not null unique comment 'ロール',
  created_at timestamp not null default current_timestamp comment '作成日時',
  updated_at timestamp not null default current_timestamp on update current_timestamp comment '更新日時'
);
-- +migrate Down
drop table if exists users;
