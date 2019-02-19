USE test;

-- show processlist;

-- CREATE TABLE IF NOT EXISTS `test`.`user` (
--  `user_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
--  `user_name` VARCHAR(45) NOT NULL COMMENT '用户名称',
--  `user_age` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄',
--  `user_sex` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
--  PRIMARY KEY (`user_id`))
--  ENGINE = InnoDB
--  AUTO_INCREMENT = 1
--  DEFAULT CHARACTER SET = utf8mb4
--  COLLATE = utf8mb4_unicode_ci
--  COMMENT = '用户表'

--  SHOW CREATE TABLE user;

SELECT * FROM user;

-- INSERT INTO user (user_name, user_age, user_sex) VALUE ("eli", 25, 2);

