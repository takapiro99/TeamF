/**
 * @author Riku Nunokawa
 */

CREATE TABLE favorite (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  dialog_id INT UNSIGNED NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES user(`id`),
  FOREIGN KEY (`dialog_id`) REFERENCES dialog(`id`),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;