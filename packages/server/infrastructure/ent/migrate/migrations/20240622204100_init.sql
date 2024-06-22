-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(20) NOT NULL, `email` varchar(254) NOT NULL, `password` varchar(255) NOT NULL, `display_name` varchar(20) NULL, `avatar_url` varchar(255) NULL, `status` enum('Active','Withdrawn','Suspended','Inactive') NOT NULL DEFAULT "Active", `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "admin_users" table
CREATE TABLE `admin_users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(20) NOT NULL, `email` varchar(254) NOT NULL, `password` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "boards" table
CREATE TABLE `boards` (`id` bigint NOT NULL AUTO_INCREMENT, `user_id` bigint NOT NULL, `title` varchar(50) NOT NULL, `description` varchar(255) NULL, `thumbnail_url` varchar(255) NULL, `order` bigint NOT NULL DEFAULT 0, `status` enum('Public','Private','Archived','Deleted') NOT NULL DEFAULT "Public", `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `user_boards` bigint NULL, PRIMARY KEY (`id`), INDEX `boards_users_boards` (`user_boards`), UNIQUE INDEX `title` (`title`), CONSTRAINT `boards_users_boards` FOREIGN KEY (`user_boards`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "threads" table
CREATE TABLE `threads` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(50) NULL, `description` varchar(255) NULL, `thumbnail_url` varchar(255) NULL, `is_auto_generated` bool NOT NULL DEFAULT 0, `is_notify_on_comment` bool NOT NULL DEFAULT 1, `ip_address` varchar(64) NOT NULL, `status` enum('Open','Archived','Deleted') NOT NULL DEFAULT "Open", `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `board_id` bigint NOT NULL, `user_id` bigint NOT NULL, PRIMARY KEY (`id`), INDEX `threads_boards_threads` (`board_id`), INDEX `threads_users_threads` (`user_id`), CONSTRAINT `threads_boards_threads` FOREIGN KEY (`board_id`) REFERENCES `boards` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `threads_users_threads` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "comments" table
CREATE TABLE `comments` (`id` bigint NOT NULL AUTO_INCREMENT, `guest_name` varchar(20) NULL, `message` longtext NOT NULL, `ip_address` varchar(64) NOT NULL, `status` enum('Visible','Deleted') NOT NULL DEFAULT "Visible", `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `parent_comment_id` bigint NULL, `thread_id` bigint NOT NULL, `user_id` bigint NULL, PRIMARY KEY (`id`), INDEX `comments_comments_replies` (`parent_comment_id`), INDEX `comments_threads_comments` (`thread_id`), INDEX `comments_users_comments` (`user_id`), CONSTRAINT `comments_comments_replies` FOREIGN KEY (`parent_comment_id`) REFERENCES `comments` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT `comments_threads_comments` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `comments_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "comment_attachments" table
CREATE TABLE `comment_attachments` (`id` bigint NOT NULL AUTO_INCREMENT, `url` varchar(255) NOT NULL, `order` bigint NOT NULL DEFAULT 0, `type` enum('image','video') NOT NULL, `comment_id` bigint NOT NULL, PRIMARY KEY (`id`), INDEX `comment_attachments_comments_comment_attachments` (`comment_id`), CONSTRAINT `comment_attachments_comments_comment_attachments` FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "thread_tags" table
CREATE TABLE `thread_tags` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(20) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`), UNIQUE INDEX `threadtag_name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "thread_taggings" table
CREATE TABLE `thread_taggings` (`thread_id` bigint NOT NULL, `tag_id` bigint NOT NULL, PRIMARY KEY (`thread_id`, `tag_id`), INDEX `thread_taggings_thread_tags_tag` (`tag_id`), CONSTRAINT `thread_taggings_thread_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `thread_tags` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `thread_taggings_threads_thread` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_board_likes" table
CREATE TABLE `user_board_likes` (`liked_at` timestamp NOT NULL, `user_id` bigint NOT NULL, `board_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `board_id`), INDEX `user_board_likes_boards_board` (`board_id`), CONSTRAINT `user_board_likes_boards_board` FOREIGN KEY (`board_id`) REFERENCES `boards` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `user_board_likes_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_board_subscriptions" table
CREATE TABLE `user_board_subscriptions` (`is_notified` bool NOT NULL DEFAULT 1, `is_checked` bool NOT NULL DEFAULT 0, `subscribed_at` timestamp NOT NULL, `user_id` bigint NOT NULL, `board_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `board_id`), INDEX `user_board_subscriptions_boards_board` (`board_id`), CONSTRAINT `user_board_subscriptions_boards_board` FOREIGN KEY (`board_id`) REFERENCES `boards` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `user_board_subscriptions_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_comment_likes" table
CREATE TABLE `user_comment_likes` (`liked_at` timestamp NOT NULL, `user_id` bigint NOT NULL, `comment_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `comment_id`), INDEX `user_comment_likes_comments_comment` (`comment_id`), CONSTRAINT `user_comment_likes_comments_comment` FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `user_comment_likes_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_comment_subscriptions" table
CREATE TABLE `user_comment_subscriptions` (`is_notified` bool NOT NULL DEFAULT 1, `is_checked` bool NOT NULL DEFAULT 0, `subscribed_at` timestamp NOT NULL, `user_id` bigint NOT NULL, `comment_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `comment_id`), INDEX `user_comment_subscriptions_comments_comment` (`comment_id`), CONSTRAINT `user_comment_subscriptions_comments_comment` FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `user_comment_subscriptions_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_thread_likes" table
CREATE TABLE `user_thread_likes` (`liked_at` timestamp NOT NULL, `user_id` bigint NOT NULL, `thread_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `thread_id`), INDEX `user_thread_likes_threads_thread` (`thread_id`), CONSTRAINT `user_thread_likes_threads_thread` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `user_thread_likes_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_thread_subscriptions" table
CREATE TABLE `user_thread_subscriptions` (`is_notified` bool NOT NULL DEFAULT 1, `is_checked` bool NOT NULL DEFAULT 0, `subscribed_at` timestamp NOT NULL, `user_id` bigint NOT NULL, `thread_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `thread_id`), INDEX `user_thread_subscriptions_threads_thread` (`thread_id`), CONSTRAINT `user_thread_subscriptions_threads_thread` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `user_thread_subscriptions_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
