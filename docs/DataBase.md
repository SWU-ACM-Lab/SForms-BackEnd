# DataBase #

**We just support MySql Database upper than 5.7.**

Here's the design document of databases. 

> We use ${xyz} to descripe a variable which named xyz.

## Permission ##

Our system has four kind of users:

+ Admin: He has complete control over the database
+ Controler: He can manage the tables
+ User: He can add a record and update the record he added
+ Guest: He can just query the public data

## Summary ##

There are two kinds of tables in our system, static tables and dynamic tables.

+ Static tables: to storage the data based on system
+ Dynamic tables: to storage the data based on service

For example, users will be storaged in static tables, and records of forms will be storaged in dynamic tables.

Here's the descriptions of the schema and its static tables.

| Table Name |           Description            | Notes |
| :--------: | :------------------------------: | :---: |
|   Users    | the information storage of users |       |
|            |                                  |       |
|            |                                  |       |
|            |                                  |       |
|            |                                  |       |
|            |                                  |       |
|            |                                  |       |

Here's the descriptions of the type of dynamic tables.



## Table: Users ##

Hers's the strucutre of  `Users` table, it will storage the information of users:

|    Field     |      Type       |   Key   |                     Description                     |
| :----------: | :-------------: | :-----: | :-------------------------------------------------: |
|   UserName   | string(varchar) |  null   |             User's name, support emoji              |
|    UserId    |       int       | primary |     User's Id, the primary key of `Users` table     |
| UserPassword | string(varchar) |  null   | User's password, which has encoded by MD5 with salt |
|  UserToken   | string(varchar) |   key   |       User's token, which used in api service       |
|  UserStatus  |       int       |  null   |         If the user is blocked or logouted          |

We will initialize the `Users` table by using this sql scripts in the database:

```sql
CREATE TABLE `${database_name}`.`Users`  (
  `UserName` varbinary(255) NOT NULL,
  `UserId` int(0) ZEROFILL NOT NULL AUTO_INCREMENT,
  `UserPassword` varchar(255) NOT NULL,
  `UserToken` varchar(255) NOT NULL,
 	`UserStatus` int(0) NOT NULL,
  PRIMARY KEY (`UserId`, `UserToken`)
);
```

## Table: Forms ##

Hers's the stucture of `Forms` table, it will storage the data of Forms which users created.

