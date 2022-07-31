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


CREATE TABLE `loanapp`.`loan_applications` (
                                               `id` INT NOT NULL,
                                               `first_name` VARCHAR(200) NULL,
                                               `middle_name` VARCHAR(200) NULL,
                                               `sur_name` VARCHAR(200) NULL,
                                               `birthday` VARCHAR(45) NULL,
                                               `pan_number` VARCHAR(45) NULL,
                                               `gender` VARCHAR(45) NULL,
                                               `user_id` INT NULL,
                                               `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                               PRIMARY KEY (`id`));


ALTER TABLE `loanapp`.`loan_applications`
    ADD COLUMN `loan_number` VARCHAR(200) NULL AFTER `created_at`,
ADD COLUMN `disbursement_id` INT NULL AFTER `loan_number`,
ADD UNIQUE INDEX `loan_number_UNIQUE` (`loan_number` ASC) VISIBLE;
;

ALTER TABLE `loanapp`.`loan_applications`
    ADD COLUMN `pancard_image` VARCHAR(200) NULL AFTER `disbursement_id`;

ALTER TABLE `loanapp`.`loan_applications`
    CHANGE COLUMN `id` `id` INT NOT NULL AUTO_INCREMENT ;

