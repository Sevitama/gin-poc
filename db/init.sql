CREATE TABLE accounts (
  id INTEGER PRIMARY KEY,
  username VARCHAR(30) NOT NULL,
  passwordhash VARCHAR(100) NOT NULL
);

CREATE TABLE article (
  id INTEGER PRIMARY KEY,
  title VARCHAR(30) NOT NULL,
  content VARCHAR(200) NOT NULL
);

INSERT INTO accounts (id, username, passwordhash) VALUES
(1001, 'mmartinez', '$2a$10$bcY3nIkDGlWizbBlb8By1ey7QeT1vlNjjcsNxt0lLwAr3XpUpSXry'),
(1002, 'r1ehrens', 'apfel'),
(1003, 's1grimm', 'biere');

INSERT INTO article (id, title, content) VALUES
(1001, 'Voting result', 'All votings are a draft.'),
(1002, 'Teslas new invention', 'The newest Tesla has now legs instead of wheels.'),
(1003, 'Security in software analysed', 'Around 10% of software engineers do use input validation and prepared statements. The rest do whatever they want.');
(1004, '<script>alert("hi")</script>', 'Around 10% of software engineers do use input validation and prepared statements. The rest do whatever they want.');
