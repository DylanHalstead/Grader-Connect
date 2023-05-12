CREATE TABLE assignment_grader (
  assignment_id INTEGER NOT NULL REFERENCES assignment(id),
  grader_id INTEGER NOT NULL REFERENCES grader(id),
  is_graded BOOLEAN NOT NULL,
  PRIMARY KEY (assignment_id, grader_id)
);
