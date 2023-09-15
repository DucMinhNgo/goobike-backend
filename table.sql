CREATE TABLE todo_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    description TEXT,
    status ENUM('doing', 'done', 'deleted') DEFAULT 'doing',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY, -- Cột này có thể sử dụng auto-increment nếu bạn muốn có một trường ID duy nhất cho mỗi người dùng
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- Set Index (Tìm kiếm thứ cấp)
-- Thay đổi dữ liệu cập nhật Index
ALTER TABLE todo_items
ADD INDEX index_status (status);

-- Primary key có hỗ trợ index (đánh 2 cột làm primary key: cột đầu sẽ được xem như là primary và có thứ tự)

EXPLAIN
SELECT *
FROM todo_items
WHERE status = 'doing'
-- WHERE id = 1
