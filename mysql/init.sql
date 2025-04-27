USE goboard;

CREATE USER IF NOT EXISTS 'goboard-user'@'%' IDENTIFIED BY 'goboard-pass';

GRANT ALL PRIVILEGES ON `goboard`.* TO 'goboard-user'@'%';

FLUSH PRIVILEGES;