create table if not exists ideas (
	id integer unsigned auto_increment primary key,
	public_id varchar(100) not null,
	user_id integer not null,
	title varchar(100),
	contents text,
	written_at datetime,
	created_at datetime
);

create table if not exists users (
	id integer unsigned auto_increment primary key,
	public_id integer unsigned not null,
	name varchar(100) not null,
  email varchar(100) not null,
	created_at datetime
);

insert into ideas (public_id, user_id, title, contents, written_at, created_at) values
	('704d9697-4c5e-c548-059e-3d67bd90e070', 1, 'first idea', 'This is first idea', now(), now());

insert into ideas (public_id, user_id, title, contents, written_at, created_at) values
	('a52bc272-71aa-2017-4d0a-f30d21b56c2a', 2, 'second idea', 'This is second idea', now(), now());

insert into users (public_id, name, email, created_at) values
	('80898b68-86c7-13bb-9fac-2a910dec44db', 'first user', 'abc@gmail.com', now());

insert into users (public_id, name, email, created_at) values
	('a8f92b26-a7fb-e849-11bd-c2692b1650c5', 'second user', 'def@gmail.com', now());

