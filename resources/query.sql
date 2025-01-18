--  ======================= TABLE SCHOOLS ============================ --

-- name: ListSchool :many
SELECT * FROM tbl_schools
ORDER BY name;

-- name: ListSchoolOnlyShortCodes :many
SELECT short_code FROM tbl_schools
ORDER BY created_at;

-- name: GetSchoolById :one
SELECT * FROM tbl_schools
WHERE id = sqlc.arg(school_id) LIMIT 1;

-- name: SearchSchoolByName :many
SELECT * FROM tbl_schools
WHERE name LIKE ?;

-- name: RegisterSchool :execresult
INSERT INTO tbl_schools (
    name,
    short_code,
    school_national_id,
    address,
    latitude,
    longitude,
    headmaster_name,
    headmaster_id,
    website,
    email,
    app_name
) VALUES (
    sqlc.arg(school_name),
    sqlc.arg(short_code),
    sqlc.arg(school_national_id),
    sqlc.narg(address),
    sqlc.narg(latitude),
    sqlc.narg(longitude),
    sqlc.arg(headmaster_name),
    sqlc.arg(headmaster_id),
    sqlc.narg(website),
    sqlc.narg(email),
    sqlc.arg(app_name)
);

-- name: UpdateSchoolById :exec
UPDATE tbl_schools
SET
    name = CASE WHEN sqlc.narg(new_name) IS NOT NULL
        THEN sqlc.narg(new_name)
        ELSE name
    END,
    
    school_national_id = CASE WHEN sqlc.narg(new_school_national_id) IS NOT NULL
        THEN sqlc.narg(new_school_national_id)
        ELSE school_national_id
    END,

    address = CASE WHEN sqlc.narg(new_address) IS NOT NULL
        THEN sqlc.narg(new_address)
        ELSE address
    END,

    latitude = CASE WHEN sqlc.narg(new_latitude) IS NOT NULL
        THEN sqlc.narg(new_latitude)
        ELSE latitude
    END,

    longitude = CASE WHEN sqlc.narg(new_longitude) IS NOT NULL
        THEN sqlc.narg(new_longitude)
        ELSE longitude
    END,

    headmaster_name = CASE WHEN sqlc.narg(new_headmaster_name) IS NOT NULL
        THEN sqlc.narg(new_headmaster_name)
        ELSE headmaster_name
    END,

    headmaster_id = CASE WHEN sqlc.narg(new_headmaster_id) IS NOT NULL
        THEN sqlc.narg(new_headmaster_id)
        ELSE headmaster_id
    END,

    website = CASE WHEN sqlc.narg(new_website) IS NOT NULL
        THEN sqlc.narg(new_website)
        ELSE website
    END,

    email = CASE WHEN sqlc.narg(new_email) IS NOT NULL
        THEN sqlc.narg(new_email)
        ELSE email
    END,

    logo_url = CASE WHEN sqlc.narg(new_logo_url) IS NOT NULL
        THEN sqlc.narg(new_logo_url)
        ELSE logo_url
    END,

    app_name = CASE WHEN sqlc.narg(new_app_name) IS NOT NULL
        THEN sqlc.narg(new_app_name)
        ELSE app_name
    END,

    updated_at = NOW()
WHERE id = ?;

--  ======================= END TABLE SCHOOLS ============================ --


--  ======================= TABLE SEMESTERS ============================ --

-- name: GetSemesterById :one
SELECT sqlc.embed(s), sqlc.embed(ss) FROM tbl_semesters as s
INNER JOIN tbl_schools as ss ON ss.id = s.school_id
WHERE s.id = ? AND s.school_id = ?
LIMIT 1;

-- name: GetActiveSemester :one
SELECT sqlc.embed(s), sqlc.embed(ss) FROM tbl_semesters as s
INNER JOIN tbl_schools as ss ON ss.id = s.school_id
WHERE s.is_active = 1 AND s.school_id = ?
LIMIT 1;

-- name: ListSemester :many
SELECT sqlc.embed(s), sqlc.embed(ss) FROM tbl_semesters as s
INNER JOIN tbl_schools as ss ON ss.id = s.school_id
WHERE s.school_id = ? AND s.is_active = ?
ORDER BY ss.name;

-- name: UpdateActiveSemester :exec
UPDATE tbl_semesters SET is_active = 1
WHERE id = ? AND school_id = ?;

-- name: DeactiveCurrentSemester :exec
UPDATE tbl_semesters SET is_active = 0
WHERE is_active = 1 AND school_id = ?;

-- name: DeleteSemesterById :execresult
DELETE FROM tbl_semesters as s
WHERE s.id = ? AND s.school_id = ? AND s.is_active = 0;

--  ======================= END TABLE SEMESTERS ============================ --



--  ======================= TABLE MAJORS ============================

-- name: ListMajor :many
SELECT * FROM tbl_majors
WHERE school_id = ?
ORDER BY name;

-- name: GetMajorByShortCode :one
SELECT * FROM tbl_majors
WHERE school_id = ? AND short_code = ?
LIMIT 1;

-- name: GetMajorById :one
SELECT * FROM tbl_majors
WHERE id = ?
LIMIT 1;

-- name: RegisterMajor :execresult
INSERT INTO tbl_majors (
    name,
    short_code,
    school_id
) VALUES (
    sqlc.arg(name),
    sqlc.arg(short_code),
    sqlc.arg(school_id)
);

-- name: UpdateMajorById :exec
UPDATE tbl_majors
SET name = CASE WHEN sqlc.narg(new_name) IS NOT NULL
    THEN sqlc.narg(new_name)
    ELSE name
END,
    short_code = CASE WHEN sqlc.narg(new_short_code) IS NOT NULL
        THEN sqlc.narg(new_short_code)
        ELSE short_code
END
WHERE id = ?;

-- name: UpdateMajorByShortCode :exec
UPDATE tbl_majors
SET name = CASE WHEN sqlc.narg(new_name) IS NOT NULL
    THEN sqlc.narg(new_name)
    ELSE name
END,
    short_code = CASE WHEN sqlc.narg(new_short_code) IS NOT NULL
        THEN sqlc.narg(new_short_code)
        ELSE short_code
END
WHERE short_code = ? AND school_id = ?;

-- name: DeleteMajorById :exec
DELETE FROM tbl_majors WHERE id = ?;

--  ======================= END TABLE MAJORS ============================



--  ======================= TABLE CLASSES ============================

-- name: ListClass :many
SELECT * FROM tbl_classes
WHERE semester_id = ? AND school_id = ?
ORDER BY name;

-- name: GetClassByName :one
SELECT * FROM tbl_classes
WHERE name LIKE ?
LIMIT 1;

-- name: ListClassByGradeLevel :many
SELECT * FROM tbl_classes
WHERE grade_level = ?
    AND semester_id = ?
    AND school_id = ?
ORDER BY name;

-- name: ListClassByMajorId :many
SELECT * FROM tbl_classes
WHERE semester_id = ? AND school_id = ? AND major_id = ?
ORDER BY name;

-- name: RegisterClass :execresult
INSERT INTO tbl_classes (
    name,
    grade_level,
    students_count,
    major_id,
    semester_id,
    school_id
) VALUES (
    sqlc.arg(name),
    sqlc.arg(grade_level),
    0,
    sqlc.arg(major_id),
    sqlc.arg(semester_id),
    sqlc.arg(school_id)
);

-- name: UpdateClass :exec
UPDATE tbl_classes
SET name = CASE WHEN sqlc.narg(new_name) IS NOT NULL
    THEN sqlc.narg(new_name)
    ELSE name
END,
    grade_level = CASE WHEN sqlc.narg(new_level) IS NOT NULL
    THEN sqlc.narg(new_level)
    ELSE grade_level
END,
    major_id = CASE WHEN sqlc.narg(new_major_id) IS NOT NULL
    THEN sqlc.narg(new_major_id)
    ELSE major_id
END,
    homeroom_teacher_id = CASE WHEN sqlc.narg(new_homeroom_teacher_id) IS NOT NULL
    THEN sqlc.narg(new_homeroom_teacher_id)
    ELSE homeroom_teacher_id
END,
    updated_at = NOW()
WHERE semester_id = ? AND school_id = ?;

--  ======================= END TABLE CLASSES ============================


--  ======================= TABLE SUBJECTS ============================

-- name: ListSubject :many
SELECT * FROM tbl_subjects
WHERE semester_id = ? AND school_id = ?
ORDER BY name;

-- name: ListSubjectByMajor :many
SELECT * FROM tbl_subjects
WHERE semester_id = ? AND school_id = ? AND major_id = ?
ORDER BY name;

-- name: GetSubjectById :one
SELECT sqlc.embed(t), sqlc.embed(sta) FROM tbl_subjects as t
LEFT JOIN tbl_subjects_access as sa
    ON sa.subject_id = t.id
LEFT JOIN tbl_subjects_teacher_access as sta
        ON sta.subject_id = t.id
WHERE t.id = ?
LIMIT 1;

-- name: GetSubjectByShortCode :one
SELECT sqlc.embed(t), sqlc.embed(sta) FROM tbl_subjects as t
LEFT JOIN tbl_subjects_access as sa
    ON sa.subject_id = t.id
LEFT JOIN tbl_subjects_teacher_access as sta
        ON sta.subject_id = t.id
WHERE t.short_code = ? AND t.semester_id = ? AND t.school_id = ?
LIMIT 1;

-- name: ListActiveSubject :many
SELECT * FROM tbl_subjects
WHERE semester_id = ? AND school_id = ? AND is_active = 1
ORDER BY name;

-- name: ListUniversalSubject :many
SELECT * FROM tbl_subjects
WHERE semester_id = ? AND school_id = ? AND is_universal = 1
ORDER BY name;

-- RegisterSubject :execresult
INSERT INTO tbl_subjects (
    name,
    short_code,
    is_active,
    is_universal,
    major_id,
    school_id,
    semester_id
) VALUES (
    sqlc.arg(name),
    sqlc.arg(short_code),
    sqlc.narg(is_active),
    sqlc.narg(is_universal),
    sqlc.arg(major_id),
    sqlc.arg(school_id),
    sqlc.arg(semester_id)
);

-- name: AssignSubjectAccess :exec
INSERT INTO tbl_subjects_access (
    school_id,
    class_id,
    subject_id
) VALUES (
    sqlc.arg(school_id),
    sqlc.arg(class_id),
    sqlc.arg(subject_id)
);

-- name: DeassignSubjectAccess :exec
DELETE FROM tbl_subjects_access
WHERE class_id = ?
    AND subject_id = ?
    AND school_id = ?;

-- name: AssignTeacherToSubject :exec
INSERT INTO tbl_subjects_teacher_access (
    school_id,
    teacher_id,
    subject_id
) VALUES (
    sqlc.arg(school_id),
    sqlc.arg(teacher_id),
    sqlc.arg(subject_id)
);

-- name: DeassignTeacherFromSubject :exec
DELETE FROM tbl_subjects_teacher_access
WHERE
    teacher_id = ?
AND
    subject_id = ?
AND
    school_id = ?;

--  ======================= END TABLE SUBJECTS ============================


--  ======================== TABLE USERS ============================

-- name: ListUser :many
SELECT sqlc.embed(user), sqlc.embed(student), sqlc.embed(teacher), sqlc.embed(user_access)
FROM tbl_users as user
LEFT JOIN tbl_users_students AS student
    ON student.id = user.user_student_id
LEFT JOIN tbl_users_teachers AS teacher
    ON teacher.id = user.user_teacher_id
LEFT JOIN tbl_users_access AS user_access
    ON user_access.id = user.user_access_id
WHERE user.school_id = ?
ORDER BY user.id;

-- name: GetUserById :one
SELECT sqlc.embed(user), sqlc.embed(student), sqlc.embed(teacher), sqlc.embed(user_access)
FROM tbl_users as user
LEFT JOIN tbl_users_students AS student
    ON student.id = user.user_student_id
LEFT JOIN tbl_users_teachers AS teacher
    ON teacher.id = user.user_teacher_id
LEFT JOIN tbl_users_access AS user_access
    ON user_access.id = user.user_access_id
WHERE user.id = ? AND user.school_id = ?
LIMIT 1;

-- name: RegisterUser :execresult
INSERT INTO tbl_users (
    name,
    is_active,
    email,
    school_id,
    password
) VALUES (
    sqlc.arg(username),
    sqlc.narg(is_active),
    sqlc.arg(email),
    sqlc.arg(school_id),
    sqlc.arg(password)
);

-- name: UpdateUser :exec
UPDATE tbl_users
SET
    name = CASE WHEN sqlc.narg(new_name) IS NOT NULL
        THEN sqlc.narg(new_name)
        ELSE name
    END,

    email = CASE WHEN sqlc.narg(new_email) IS NOT NULL
        THEN sqlc.narg(new_email)
        ELSE email
    END,

    password = CASE WHEN sqlc.narg(new_password) IS NOT NULL
        THEN sqlc.narg(new_password)
        ELSE password
    END,

    user_access_id = CASE WHEN sqlc.narg(new_user_access_id) IS NOT NULL
        THEN sqlc.narg(new_user_access_id)
        ELSE user_access_id
    END,

    user_student_id = CASE WHEN sqlc.narg(new_user_student_id) IS NOT NULL
        THEN sqlc.narg(new_user_student_id)
        ELSE user_student_id
    END,

    user_teacher_id = CASE WHEN sqlc.narg(new_user_teacher_id) IS NOT NULL
        THEN sqlc.narg(new_user_teacher_id)
        ELSE user_teacher_id
    END,
    
    updated_at = NOW()
WHERE id = ?;

-- name: SearchUser :many
SELECT * FROM tbl_users as user
WHERE user.name LIKE sqlc.arg(query)
    OR user.email LIKE sqlc.arg(query)
    AND user.school_id = sqlc.arg(school_id)
ORDER BY user.id;

-- name: GetUserByIdAndAccess :one
SELECT sqlc.embed(user), sqlc.embed(user_access) FROM tbl_users as user
INNER JOIN tbl_users_access as user_access
    ON user.id = user_access.user_id
WHERE user.id = sqlc.arg(user_id)
    AND user_access.is_student = sqlc.arg(is_student)
    AND user_access.is_admin = sqlc.arg(is_admin)
    AND user_access.is_teacher = sqlc.arg(is_teacher)
LIMIT 1;

-- name: GetUserByNameOrEmail :one
SELECT sqlc.embed(user), sqlc.embed(user_access) FROM tbl_users as user
INNER JOIN tbl_users_access as user_access
    ON user.id = user_access.user_id
WHERE (user.email = sqlc.arg(user_query)
    OR user.name = sqlc.arg(user_query))
    AND user.school_id = sqlc.arg(school_id)
LIMIT 1;

-- name: CreateUserAccess :execresult
INSERT INTO tbl_users_access (
    user_id,
    school_id,
    is_student,
    is_admin,
    is_teacher
) VALUES (
    sqlc.arg(user_id),
    sqlc.arg(school_id),
    sqlc.arg(is_student),
    sqlc.arg(is_admin),
    sqlc.arg(is_teacher)
);

-- name: UpdateUserAccess :exec
UPDATE tbl_users_access
SET
    is_student = CASE WHEN sqlc.narg(is_student) IS NOT NULL
        THEN sqlc.narg(is_student)
        ELSE is_student
    END,

    is_admin = CASE WHEN sqlc.narg(is_admin) IS NOT NULL
        THEN sqlc.narg(is_admin)
        ELSE is_admin
    END,

    is_teacher = CASE WHEN sqlc.narg(is_teacher) IS NOT NULL
        THEN sqlc.narg(is_teacher)
        ELSE is_teacher
    END
WHERE id = ?;

-- name: DeleteUserAccess :exec
DELETE FROM tbl_users_access WHERE id = ?;

-- name: DeleteUserAccessByUserId :exec
DELETE FROM tbl_users_access WHERE user_id = ?;

--  ======================= END TABLE USERS ============================



--  ======================= TABLE STUDENTS ============================

-- name: ListStudent :many
SELECT sqlc.embed(student), sqlc.embed(user), sqlc.embed(class)
FROM tbl_users_students as student
LEFT JOIN tbl_users AS user
    ON user.user_student_id = student.id
LEFT JOIN tbl_classes AS class
    ON class.id = student.class_id
WHERE student.school_id = sqlc.arg(school_id)
    AND student.semester_id = sqlc.arg(semester_id)
ORDER BY student.id;


-- name: GetStudentById :one
SELECT sqlc.embed(student), sqlc.embed(user), sqlc.embed(class)
FROM tbl_users_students as student
LEFT JOIN tbl_users AS user
    ON user.user_student_id = student.id
LEFT JOIN tbl_classes AS class
    ON class.id = student.class_id
WHERE student.school_id = sqlc.arg(school_id)
    AND student.semester_id = sqlc.arg(semester_id)
    AND student.id = sqlc.arg(id)
LIMIT 1;

-- name: SearchStudent :many
SELECT sqlc.embed(student), sqlc.embed(user), sqlc.embed(class)
FROM tbl_users_students as student
LEFT JOIN tbl_users AS user
    ON user.user_student_id = student.id
LEFT JOIN tbl_classes AS class
    ON class.id = student.class_id
WHERE student.school_id = sqlc.arg(school_id)
    AND student.semester_id = sqlc.arg(semester_id)
    AND (
        student.name LIKE sqlc.arg(query)
        OR student.national_student_id LIKE sqlc.arg(query)
        OR student.school_student_id LIKE sqlc.arg(query)
    )
ORDER BY student.id;

-- name: RegisterStudent :execresult
INSERT INTO tbl_users_students (
    name,
    national_student_id,
    school_student_id,
    religion,

    user_id,
    school_id,
    semester_id,
    class_id
) VALUES (
    sqlc.arg(name),
    sqlc.arg(national_student_id),
    sqlc.arg(school_student_id),
    sqlc.arg(religion),
    sqlc.arg(user_id),
    sqlc.arg(school_id),
    sqlc.arg(semester_id),
    sqlc.narg(class_id)
);

-- name: UpdateStudent :exec
UPDATE tbl_users_students
SET
    name = CASE WHEN sqlc.narg(new_name) IS NOT NULL
        THEN sqlc.narg(new_name)
        ELSE name
    END,

    national_student_id = CASE WHEN sqlc.narg(new_national_student_id) IS NOT NULL
        THEN sqlc.narg(new_national_student_id)
        ELSE national_student_id
    END,

    school_student_id = CASE WHEN sqlc.narg(new_school_student_id) IS NOT NULL
        THEN sqlc.narg(new_school_student_id)
        ELSE school_student_id
    END,

    religion = CASE WHEN sqlc.narg(new_religion) IS NOT NULL
        THEN sqlc.narg(new_religion)
        ELSE religion
    END,

    class_id = CASE WHEN sqlc.narg(new_class_id) IS NOT NULL
        THEN sqlc.narg(new_class_id)
        ELSE class_id
    END,

    updated_at = NOW()
WHERE id = ?;

-- name: DeleteStudent :exec
DELETE FROM tbl_users_students
WHERE id = ?;

--  ======================= END TABLE STUDENTS ============================


--  ======================= TABLE TEACHERS ============================

-- name: ListTeacher :many
SELECT sqlc.embed(teacher), sqlc.embed(user)
FROM tbl_users_teachers AS teacher
LEFT JOIN tbl_users AS user
    ON user.id = teacher.user_id
WHERE teacher.semester_id = sqlc.arg(semester_id)
    AND teacher.school_id = sqlc.arg(school_id)
ORDER BY teacher.id;

-- name: GetTeacherById :one
SELECT sqlc.embed(teacher), sqlc.embed(user)
FROM tbl_users_teachers AS teacher
LEFT JOIN tbl_users AS user
    ON user.id = teacher.user_id
WHERE teacher.semester_id = sqlc.arg(semester_id)
    AND teacher.school_id = sqlc.arg(school_id)
    AND teacher.id = sqlc.arg(id)
LIMIT 1;

-- name: SearchTeacher :many
SELECT *
FROM tbl_users_teachers
WHERE name LIKE sqlc.arg(query)
    OR national_teacher_id LIKE sqlc.arg(query)
    AND school_id = sqlc.arg(school_id)
    AND semester_id = sqlc.arg(semester_id)
ORDER BY teacher.id;

-- name: RegisterTeacher :execresult
INSERT INTO tbl_users_teachers (
    name,
    national_teacher_id,
    school_id,
    user_id,
    semester_id
) VALUES (
    sqlc.arg(name),
    sqlc.arg(national_teacher_id),
    sqlc.arg(school_id),
    sqlc.arg(user_id),
    sqlc.arg(semester_id)
);

-- name: UpdateTeacher :exec
UPDATE tbl_users_teachers
SET
    name = CASE WHEN sqlc.narg(new_name) IS NOT NULL
        THEN sqlc.narg(new_name)
        ELSE name
    END,

    national_teacher_id = CASE WHEN sqlc.narg(new_national_teacher_id) IS NOT NULL
        THEN sqlc.narg(new_national_teacher_id)
        ELSE national_teacher_id
    END,

    updated_at = NOW()
WHERE id = ?;

-- name: DeleteTeacher :exec
DELETE FROM tbl_users_teachers
WHERE id = ?;

--  ======================= END TABLE TEACHERS ============================



--  ======================= TABLE AUTH LOGS ============================

-- name: ListAuthLog :many
SELECT *
FROM tbl_auth_logs
WHERE school_id = ?
ORDER BY id;

-- name: CreateAuthLog :execresult
INSERT INTO tbl_auth_logs (
    user_id,
    client_ip,
    client_useragent,
    school_id
) VALUES (
    sqlc.arg(user_id),
    sqlc.arg(client_ip),
    sqlc.arg(client_useragent),
    sqlc.arg(school_id)
);


--  ======================= END TABLE AUTH LOGS ============================
