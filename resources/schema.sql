SET FOREIGN_KEY_CHECKS = 0;

-- ======================= SCHOOL MANAGEMENT ============================ --

CREATE TABLE tbl_schools (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    short_code VARCHAR(255) NOT NULL,
    school_national_id VARCHAR(20) NOT NULL,
    is_active INT(1) DEFAULT 1,
    address TEXT,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    headmaster_name VARCHAR(255) NOT NULL,
    headmaster_id VARCHAR(255) NOT NULL,
    website VARCHAR(255),
    email VARCHAR(255),
    logo_url VARCHAR(255),
    app_name VARCHAR(255) NOT NULL DEFAULT "GarudaCBTX",
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
) ENGINE=InnoDB;

CREATE TABLE tbl_semesters (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cycle INT(1) NOT NULL,
    year INT NOT NULL,
    is_active INT(1) DEFAULT 0,
    school_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE tbl_majors (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    short_code VARCHAR(255) NOT NULL,
    school_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE tbl_classes (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    grade_level INT(10) NOT NULL,
    students_count INT NOT NULL,
    major_id BIGINT NOT NULL,
    semester_id BIGINT NOT NULL,
    school_id BIGINT NOT NULL,
    homeroom_teacher_id BIGINT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (major_id) REFERENCES tbl_majors(id) ON DELETE CASCADE,
    FOREIGN KEY (semester_id) REFERENCES tbl_semesters(id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE,
    FOREIGN KEY (homeroom_teacher_id) REFERENCES tbl_users_teachers(id) ON DELETE SET NULL
) ENGINE=InnoDB;

-- ======================= ASSIGNMENT MANAGEMENT ============================ --

CREATE TABLE tbl_subjects (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    short_code VARCHAR(255) NOT NULL,
    is_active INT(1) DEFAULT 1,
    is_universal INT(1) NOT NULL DEFAULT 1,
    major_id BIGINT NOT NULL,
    school_id BIGINT NOT NULL,
    semester_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (major_id) REFERENCES tbl_majors(id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE,
    FOREIGN KEY (semester_id) REFERENCES tbl_semesters(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE tbl_subjects_access (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    school_id BIGINT NOT NULL,
    class_id BIGINT NOT NULL,
    subject_id BIGINT NOT NULL,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE,
    FOREIGN KEY (class_id) REFERENCES tbl_classes(id) ON DELETE CASCADE,
    FOREIGN KEY (subject_id) REFERENCES tbl_subjects(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE tbl_subjects_teacher_access (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    subject_id BIGINT NOT NULL,
    school_id BIGINT NOT NULL,
    teacher_id BIGINT NOT NULL,
    FOREIGN KEY (subject_id) REFERENCES tbl_subjects(id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE,
    FOREIGN KEY (teacher_id) REFERENCES tbl_users_teachers(id) ON DELETE CASCADE
) ENGINE=InnoDB;

-- ======================= USERS MANAGEMENT ============================ --

CREATE TABLE tbl_users (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    is_active INT(1) DEFAULT 0,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    user_access_id BIGINT NULL,
    user_student_id BIGINT NULL,
    user_teacher_id BIGINT NULL,
    school_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_access_id) REFERENCES tbl_users_access(id) ON DELETE SET NULL,
    FOREIGN KEY (user_student_id) REFERENCES tbl_users_students(id) ON DELETE SET NULL,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE,
    FOREIGN KEY (user_teacher_id) REFERENCES tbl_users_teachers(id) ON DELETE SET NULL
) ENGINE=InnoDB;

CREATE TABLE tbl_users_access (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    school_id BIGINT NOT NULL,
    is_student INT(1) NOT NULL,
    is_admin INT(1) NOT NULL,
    is_teacher INT(1) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES tbl_users(id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE tbl_users_teachers (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    national_teacher_id VARCHAR(255) NOT NULL UNIQUE,
    school_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    semester_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES tbl_users(id) ON DELETE CASCADE,
    FOREIGN KEY (semester_id) REFERENCES tbl_semesters(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE tbl_users_students (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    national_student_id VARCHAR(20) NOT NULL,
    school_student_id VARCHAR(20) NOT NULL,
    religion ENUM('islam', 'kristen', 'katolik', 'konghucu', 'buddha', 'hindu') NOT NULL,
    user_id BIGINT NOT NULL,
    school_id BIGINT NOT NULL,
    class_id BIGINT NULL,
    semester_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES tbl_users(id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE,
    FOREIGN KEY (class_id) REFERENCES tbl_classes(id) ON DELETE SET NULL,
    FOREIGN KEY (semester_id) REFERENCES tbl_semesters(id) ON DELETE CASCADE
) ENGINE=InnoDB;

-- ======================= AUTH MANAGEMENT ============================ --

CREATE TABLE tbl_auth_logs (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    client_ip VARCHAR(25) NOT NULL,
    client_useragent TEXT NOT NULL,
    school_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES tbl_users(id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES tbl_schools(id) ON DELETE CASCADE
) ENGINE=InnoDB;

SET FOREIGN_KEY_CHECKS = 1;
