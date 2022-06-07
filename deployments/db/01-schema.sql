DROP TABLE IF EXISTS top100stats;

CREATE TABLE top100stats (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   criterion VARCHAR(32) NOT NULL,
   rank SMALLINT NOT NULL,
   chain VARCHAR(32) NOT NULL,
   nft_id VARCHAR(128) NOT NULL,
   nft_slug VARCHAR(128) NOT NULL,
   nft_name VARCHAR(128) NOT NULL,
   daily_volume decimal NOT NULL,
   volume_delta decimal NOT NULL,
   PRIMARY KEY (id)
 );
