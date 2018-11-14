/***************************************************
	MIGRATION down
	Id: 72922af6-a5bd-4234-a636-3779004a9a4f
	Description: add-friend
***************************************************/

CREATE TABLE friend (
  id int,
  firstName varchar(255),
  lastName varchar(255),
  address varchar(255),
  city varchar(255)
);