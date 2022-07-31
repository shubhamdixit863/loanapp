CREATE TABLE `loanapp`.`users` (
                                   `id` INT NOT NULL AUTO_INCREMENT,
                                   `first_name` VARCHAR(200) NULL,
                                   `last_name` VARCHAR(200) NULL,
                                   `email` VARCHAR(200) NULL,
                                   `password` VARCHAR(250) NULL,
                                   `phone` VARCHAR(45) NULL,
                                   PRIMARY KEY (`id`),
                                   UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
                                   UNIQUE INDEX `phone_UNIQUE` (`phone` ASC) VISIBLE);
