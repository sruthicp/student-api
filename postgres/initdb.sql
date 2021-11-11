CREATE SCHEMA school;

CREATE TABLE school.student
(
    admission_no integer NOT NULL,
    std_name character varying(50) NOT NULL,
    std_address character varying(50) NOT NULL,
    std_class character varying(10) NOT NULL,
    std_age integer NOT NULL,
    CONSTRAINT student_key UNIQUE (admission_no)
);

