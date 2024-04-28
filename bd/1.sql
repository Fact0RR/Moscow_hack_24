CREATE TABLE vacancies(
  id serial PRIMARY KEY,
  link text not null,
  CONSTRAINT link_vacancies_unique UNIQUE (link)
);

CREATE TABLE courses(
  id serial PRIMARY KEY,
  link text not null,
  name text not null,
  type text not null,
  duration text not null,

  CONSTRAINT link_courses_unique UNIQUE (link),
  CONSTRAINT name_courses_unique UNIQUE (name)
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
  points int NOT NULL,
  description text
);

CREATE TABLE TAGS (
  id serial PRIMARY KEY,
  name text NOT NULL,
  description text,

  CONSTRAINT name_tags_unique UNIQUE (name)
);

CREATE TABLE C_T(
  id serial PRIMARY KEY,
  course_id int NOT NULL,
  tag_id int NOT NULL,

  CONSTRAINT unique_bid_tid UNIQUE (course_id, tag_id),

  FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES TAGS (id) ON DELETE CASCADE ON UPDATE CASCADE
);