package collection

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Persist(db *sqlx.DB, day string, nft *NFT, stats *Stats) error {
	if nft == nil {
		err := errors.New("nil nft")
		log.Error(err)
		return err
	}
	if stats == nil {
		err := errors.New("nil stats")
		log.Error(err)
		return err
	}
	err := persistNFT(db, *nft)
	if err != nil {
		return err
	}
	err = persistStats(db, day, *nft, *stats)
	if err != nil {
		return err
	}

	return nil
}

func persistNFT(db *sqlx.DB, nft NFT) error {
	log.Infof("%v", nft)
	q := `
		INSERT IGNORE INTO nft(
			data_source_name, chain, slug, name, image_url, discord_url, url,
			safelist_status)
		VALUES(
			:data_source_name, :chain, :slug, :name, :image_url, :discord_url,
			:url, :safelist_status)
		`
	_, err := db.NamedExec(q, nft)
	if err != nil {
		err = fmt.Errorf("failed to insert nft '%s', %w", nft.Slug, err)
		log.Errorf(err.Error())
		return err
	}
	return nil
}

func persistStats(db *sqlx.DB, day string, nft NFT, stats Stats) error {
	qt := `
		INSERT INTO stats(
			day, d1_volume, d1_change, d1_sales, d1_avg_price,
			d7_volume, d7_change, d7_sales, d7_avg_price,
			d30_volume, d30_change, d30_sales, d30_avg_price,
			total_volume, total_sales, total_supply,
			owners, avg_price, market_cap, floor_price, slug,
			data_source_name)
		VALUES(
			'%s', :d1_volume, :d1_change, :d1_sales, :d1_avg_price,
			:d7_volume, :d7_change, :d7_sales, :d7_avg_price,
			:d30_volume, :d30_change, :d30_sales, :d30_avg_price,
			:total_volume, :total_sales, :total_supply,
			:owners, :avg_price, :market_cap, :floor_price, '%s',
			'%s')
	`
	q := fmt.Sprintf(qt, day, nft.Slug, nft.DataSource)
	_, err := db.NamedExec(q, stats)
	if err != nil {
		err = fmt.Errorf("failed to insert stats for nft '%s', %w", nft.Slug, err)
		log.Errorf(err.Error())
		return err
	}
	return nil
}
