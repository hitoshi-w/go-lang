
-- +migrate Up
create table if not exists tasks (
  id bigint not null primary key auto_increment comment 'タスクID',
  title varchar(255) not null comment 'タイトル',
  status varchar(255) not null comment 'ステータス',
  created_at timestamp not null default current_timestamp comment '作成日時',
  updated_at timestamp not null default current_timestamp on update current_timestamp comment '更新日時'
);

-- +migrate Down
drop table if exists tasks;
