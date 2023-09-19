create schema api;
create table if not exists api.pokemon (
  id serial,
  pokedex_id smallint not null,
  name varchar(255) not null,
  type_1 varchar(15) not null,
  type_2 varchar(15) default null,
  height smallint not null,
  weight smallint not null,
  sprite_url text default 'https://archives.bulbagarden.net/media/upload/0/00/Bag_Pok%C3%A9_Ball_SV_Sprite.png',
  created_at timestamp with time zone default current_timestamp,
  modified_at timestamp with time zone default current_timestamp
);

create table if not exists api.users (
  id serial,
  email varchar(255) not null,
  pass_hash text not null,
  display_name varchar(255)
);

-- seed database
insert into api.pokemon (
  pokedex_id,
  name,
  type_1,
  type_2,
  height,
  weight,
  sprite_url
) VALUES
  (2, 'Ivysaur', 'grass', 'poison', 10, 130, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png'),
  (25, 'Pikachu', 'electric', null, 4, 60, 'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png');