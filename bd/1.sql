CREATE TABLE vacancies(
  id serial PRIMARY KEY,
  link not null,
  CONSTRAINT link_unique UNIQUE (link)
);

CREATE TABLE courses(
  id serial PRIMARY KEY,
  link not null,
  CONSTRAINT link_unique UNIQUE (link)
);

CREATE TABLE v_c(
  id serial PRIMARY KEY,
  vacancy_id int NOT NUll,
  course_id int NOT NULL,

  CONSTRAINT unique_vid_cid UNIQUE (vacancy_id, course_id),

  FOREIGN KEY (vacancy_id) REFERENCES vacancies (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE REVIEWS (
  id serial PRIMARY KEY,
  email varchar(255) NOT NULL,
  points int NOT NULL,
  description text 
);