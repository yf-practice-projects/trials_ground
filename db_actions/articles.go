package db_actions

const CreatArticles = `CREATE TABLE ARTICLES(
    ID VARCHAR(36) NOT NULL PRIMARY KEY,
    TITLE VARCHAR(180) NOT NULL,
    CONTENT VARCHAR(180) NOT NULL,
    CREATED_AT DATETIME DEFAULT NOW(),
    UPDATED_AT DATETIME DEFAULT NOW(),
    FG_DELETED BOOLEAN DEFAULT FALSE
    )`

const GetAllArticles = `SELECT
    ID
    , TITLE
    , CONTENT
    , CREATED_AT
    , UPDATED_AT 
    FROM
        ARTICLES 
    WHERE
        FG_DELETED = 0
    `

const GetArticle = `SELECT
    ID
    , TITLE
    , CONTENT
    , CREATED_AT
    , UPDATED_AT 
    FROM
        ARTICLES 
    WHERE
        FG_DELETED = 0
        AND ID = ?
    `
