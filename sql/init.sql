CREATE DATABASE poc_db;
CREATE USER poc_user WITH ENCRYPTED PASSWORD 'poc_secret';
GRANT ALL PRIVILEGES ON DATABASE poc_db TO poc_user;

\connect poc_db poc_user

CREATE TYPE "enum_RoleName" AS ENUM (
  'OWNER',
  'DIRECTOR',
  'MANAGER',
  'EMPLOYEE',
  'INTERN'
);
ALTER TYPE "enum_RoleName" OWNER TO poc_user;

CREATE TABLE "Roles"(
  id INTEGER NOT NULL,
  name "enum_RoleName" DEFAULT 'EMPLOYEE'::"enum_RoleName" UNIQUE NOT NULL,
  description TEXT
);
ALTER TABLE "Roles" OWNER TO poc_user;
CREATE SEQUENCE "Roles_id_seq"
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TABLE "Cars"(
  id INTEGER NOT NULL,
  name VARCHAR(255) UNIQUE NOT NULL,
  "registryNr" VARCHAR(255) UNIQUE NOT NULL,
  description TEXT,
  "createdAt" timestamp with time zone,
  "updatedAt" timestamp with time zone
);
ALTER TABLE "Cars" OWNER TO poc_user;
CREATE SEQUENCE "Cars_id_seq"
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TABLE "Employees"(
  id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  "RoleId" INTEGER NOT NULL,
  -- "ManagerId" INTEGER,
  "CarId" INTEGER,
  "TeamId" INTEGER,
  "hrDocumentation" jsonb,
  active boolean,
  "createdAt" timestamp with time zone,
  "updatedAt" timestamp with time zone
);
ALTER TABLE "Employees" OWNER TO poc_user;
CREATE SEQUENCE "Employees_id_seq"
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TABLE "Teams"(
  id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  "RegionId" INTEGER NOT NULL,
  "DepartmentId" INTEGER NOT NULL
);
ALTER TABLE "Teams" OWNER TO poc_user;
CREATE SEQUENCE "Teams_id_seq"
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TABLE "Departments"(
  id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  "ManagerId" INTEGER,
  hrDocumentation json,
  "createdAt" timestamp with time zone,
  "updatedAt" timestamp with time zone
);
ALTER TABLE "Departments" OWNER TO poc_user;
CREATE SEQUENCE "Departments_id_seq"
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TABLE "Regions"(
  id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  description jsonb,
  "createdAt" timestamp with time zone,
  "updatedAt" timestamp with time zone
);
ALTER TABLE "Regions" OWNER TO poc_user;
CREATE SEQUENCE "Regions_id_seq"
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

CREATE TABLE "TeamRegion"(
  "TeamId" INTEGER NOT NULL,
  "RegionId" INTEGER NOT NULL
);
ALTER TABLE "TeamRegion" OWNER TO poc_user;

ALTER TABLE ONLY "Roles" ALTER COLUMN id SET DEFAULT nextval('"Roles_id_seq"'::regClass);
ALTER TABLE ONLY "Cars" ALTER COLUMN id SET DEFAULT nextval('"Cars_id_seq"'::regClass);
ALTER TABLE ONLY "Employees" ALTER COLUMN id SET DEFAULT nextval('"Employees_id_seq"'::regClass);
ALTER TABLE ONLY "Teams" ALTER COLUMN id SET DEFAULT nextval('"Teams_id_seq"'::regClass);
ALTER TABLE ONLY "Departments" ALTER COLUMN id SET DEFAULT nextval('"Departments_id_seq"'::regClass);
ALTER TABLE ONLY "Regions" ALTER COLUMN id SET DEFAULT nextval('"Regions_id_seq"'::regClass);

ALTER TABLE ONLY "Roles" 
  ADD CONSTRAINT "Roles_pk" 
  PRIMARY KEY(id);
ALTER TABLE ONLY "Cars" 
  ADD CONSTRAINT "Cars_pk" 
  PRIMARY KEY(id);
ALTER TABLE ONLY "Employees" 
  ADD CONSTRAINT "Employees_pk" 
  PRIMARY KEY(id);
ALTER TABLE ONLY "Teams" 
  ADD CONSTRAINT "Teams_pk" 
  PRIMARY KEY(id);
ALTER TABLE ONLY "Departments" 
  ADD CONSTRAINT "Departments_pk" 
  PRIMARY KEY(id);
ALTER TABLE ONLY "Regions" 
  ADD CONSTRAINT "Regions_pk" 
  PRIMARY KEY(id);
ALTER TABLE ONLY "TeamRegion" 
  ADD CONSTRAINT "TeamRegion_pk" 
  PRIMARY KEY("TeamId","RegionId");

ALTER TABLE ONLY "Roles" 
  ADD CONSTRAINT "Roles_uq" 
  UNIQUE("name");
ALTER TABLE ONLY "Cars" 
  ADD CONSTRAINT "Cars_registryNr_uq" 
  UNIQUE("registryNr");
ALTER TABLE ONLY "Cars" 
  ADD CONSTRAINT "Cars_name_uq" 
  UNIQUE("name");

ALTER TABLE ONLY "Employees" 
  ADD CONSTRAINT "Employees_RoleId_fk"
  FOREIGN KEY("RoleId")
  REFERENCES "Roles"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
-- ALTER TABLE ONLY "Employees" 
--   ADD CONSTRAINT "Employees_ManagerId_fk"
--   FOREIGN KEY("ManagerId")
--   REFERENCES "Employees"(id)
--   ON UPDATE CASCADE
--   ON DELETE CASCADE;
ALTER TABLE ONLY "Employees" 
  ADD CONSTRAINT "Employees_CarId_fk"
  FOREIGN KEY("CarId")
  REFERENCES "Cars"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
ALTER TABLE ONLY "Employees" 
  ADD CONSTRAINT "Employees_TeamId_fk"
  FOREIGN KEY("TeamId")
  REFERENCES "Teams"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
ALTER TABLE ONLY "Departments" 
  ADD CONSTRAINT "Departments_ManagerId_fk"
  FOREIGN KEY("ManagerId")
  REFERENCES "Employees"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
ALTER TABLE ONLY "Teams" 
  ADD CONSTRAINT "Teams_DepartmentId_fk"
  FOREIGN KEY("DepartmentId")
  REFERENCES "Departments"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
ALTER TABLE ONLY "Teams" 
  ADD CONSTRAINT "Teams_RegionId_fk"
  FOREIGN KEY("RegionId")
  REFERENCES "Regions"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
ALTER TABLE ONLY "TeamRegion" 
  ADD CONSTRAINT "TeamRegion_TeamId_fk"
  FOREIGN KEY("TeamId")
  REFERENCES "Teams"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
ALTER TABLE ONLY "TeamRegion" 
  ADD CONSTRAINT "TeamRegion_RegionId_fk"
  FOREIGN KEY("TeamId")
  REFERENCES "Regions"(id)
  ON UPDATE CASCADE
  ON DELETE CASCADE;