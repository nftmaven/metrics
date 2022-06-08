DROP TABLE IF EXISTS data_source;
CREATE TABLE data_source (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   name VARCHAR(128) NOT NULL,
   uri VARCHAR(128) NOT NULL,
   api_key VARCHAR(128),
   api_secret VARCHAR(128),
   PRIMARY KEY (id)
);

INSERT INTO data_source(name, uri) VALUES('opensea', 'https://opensea.io/');


DROP TABLE IF EXISTS nft;
CREATE TABLE nft (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   date_source_id MEDIUMINT NOT NULL,
   chain VARCHAR(32) NOT NULL,
   slug VARCHAR(128) NOT NULL,
   name VARCHAR(128) NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (date_source_id) REFERENCES data_source (id)
 );


DROP TABLE IF EXISTS top100stats;
CREATE TABLE top100stats (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   day DATE NOT NULL,
   criterion VARCHAR(32) NOT NULL,
   nft_rank SMALLINT NOT NULL,
   nft_id MEDIUMINT NOT NULL,
   date_source_id MEDIUMINT NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (nft_id) REFERENCES nft (id),
   FOREIGN KEY (date_source_id) REFERENCES data_source (id)
 );


DROP TABLE IF EXISTS sales;
CREATE TABLE sales (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   day DATE NOT NULL,
   v1 DECIMAL NOT NULL,
   v1_delta DECIMAL NOT NULL,
   s1 MEDIUMINT NOT NULL,
   v7 DECIMAL NOT NULL,
   v7_delta DECIMAL NOT NULL,
   s7 MEDIUMINT NOT NULL,
   v30 DECIMAL NOT NULL,
   v30_delta DECIMAL NOT NULL,
   s30 MEDIUMINT NOT NULL,
   vt DECIMAL NOT NULL,
   st MEDIUMINT NOT NULL,
   supply MEDIUMINT NOT NULL,
   owners MEDIUMINT NOT NULL,
   avg_price DECIMAL NOT NULL,
   market_cap DECIMAL NOT NULL,
   floor_price DECIMAL NOT NULL,
   nft_id MEDIUMINT NOT NULL,
   date_source_id MEDIUMINT NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (nft_id) REFERENCES nft (id),
   FOREIGN KEY (date_source_id) REFERENCES data_source (id)
 );
