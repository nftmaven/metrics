DROP TABLE IF EXISTS data_source;
CREATE TABLE data_source (
   name VARCHAR(128) NOT NULL,
   uri VARCHAR(128) NOT NULL,
   api_key VARCHAR(128),
   api_secret VARCHAR(128),
   PRIMARY KEY (name),
   unique(uri)
);

INSERT INTO data_source(name, uri) VALUES('opensea', 'https://opensea.io/');


DROP TABLE IF EXISTS nft;
CREATE TABLE nft (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   data_source_name VARCHAR(128) NOT NULL,
   chain VARCHAR(32) NOT NULL,
   slug VARCHAR(128) NOT NULL,
   name VARCHAR(128) NOT NULL,
   image_url VARCHAR(1024) NOT NULL,
   discord_url VARCHAR(1024) NOT NULL,
   url VARCHAR(1024) NOT NULL,
   safelist_status VARCHAR(32) NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (data_source_name) REFERENCES data_source (name),
   unique(slug, data_source_name)
 );


DROP TABLE IF EXISTS top100stats;
CREATE TABLE top100stats (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   day DATE NOT NULL,
   criterion VARCHAR(32) NOT NULL,
   rank SMALLINT NOT NULL,
   slug VARCHAR(128) NOT NULL,
   data_source_name VARCHAR(128) NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (slug, data_source_name) REFERENCES nft (slug, data_source_name),
   FOREIGN KEY (data_source_name) REFERENCES data_source (name),
   unique(day, slug, data_source_name, criterion)
 );


DROP TABLE IF EXISTS stats;
CREATE TABLE stats (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   day DATE NOT NULL,

   d1_volume DECIMAL(20,8) NOT NULL,
   d1_change DECIMAL(20,8) NOT NULL,
   d1_sales DECIMAL(20,8) NOT NULL,
   d1_avg_price DECIMAL(20,8) NOT NULL,

   d7_volume DECIMAL(20,8) NOT NULL,
   d7_change DECIMAL(20,8) NOT NULL,
   d7_sales DECIMAL(20,8) NOT NULL,
   d7_avg_price DECIMAL(20,8) NOT NULL,

   d30_volume DECIMAL(20,8) NOT NULL,
   d30_change DECIMAL(20,8) NOT NULL,
   d30_sales DECIMAL(20,8) NOT NULL,
   d30_avg_price DECIMAL(20,8) NOT NULL,

   total_volume DECIMAL(20,8) NOT NULL,
   total_sales DECIMAL(20,8) NOT NULL,
   total_supply MEDIUMINT NOT NULL,

   owners MEDIUMINT NOT NULL,
   avg_price DECIMAL(20,8) NOT NULL,
   market_cap DECIMAL(20,8) NOT NULL,
   floor_price DECIMAL(20,8) NOT NULL,

   slug VARCHAR(128) NOT NULL,
   data_source_name VARCHAR(128) NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (slug, data_source_name) REFERENCES nft (slug, data_source_name),
   FOREIGN KEY (data_source_name) REFERENCES data_source (name),
   unique(day,slug, data_source_name)
 );
