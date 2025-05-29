
package shell

import (
	"time"

	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aliyun"
)

func init() {
	type MongoDBListOptions struct {
		MongoType string `choices:"sharding|replicate|serverless"`
		Offset    int
		Limit     int
	}
	shellutils.R(&MongoDBListOptions{}, "mongodb-list", "List mongodb", func(cli *aliyun.SRegion, args *MongoDBListOptions) error {
		dbs, _, err := cli.GetMongoDBs(args.MongoType, args.Limit, args.Offset)
		if err != nil {
			return err
		}
		printList(dbs, 0, 0, 0, []string{})
		return nil
	})

	type MongoDBIdOptions struct {
		ID string
	}

	shellutils.R(&MongoDBIdOptions{}, "mongodb-show", "Show mongodb", func(cli *aliyun.SRegion, args *MongoDBIdOptions) error {
		db, err := cli.GetMongoDB(args.ID)
		if err != nil {
			return errors.Wrapf(err, "GetMongoDB(%s)", args.ID)
		}
		printObject(db)
		return nil

	})

	shellutils.R(&MongoDBIdOptions{}, "mongodb-delete", "Delete mongodb", func(cli *aliyun.SRegion, args *MongoDBIdOptions) error {
		return cli.DeleteMongoDB(args.ID)
	})

	type MongoDBBackupListOptions struct {
		ID         string
		START      time.Time
		END        time.Time
		PageSize   int
		PageNumber int
	}

	shellutils.R(&MongoDBBackupListOptions{}, "mongodb-backup-list", "List mongodb backups", func(cli *aliyun.SRegion, args *MongoDBBackupListOptions) error {
		backups, _, err := cli.GetMongoDBBackups(args.ID, args.START, args.END, args.PageSize, args.PageNumber)
		if err != nil {
			return err
		}
		printList(backups, 0, 0, 0, nil)
		return nil
	})

	type MongoDBSkuListOptions struct {
	}

	shellutils.R(&MongoDBSkuListOptions{}, "mongodb-sku-list", "List mongodb skus", func(cli *aliyun.SRegion, args *MongoDBSkuListOptions) error {
		skus, err := cli.GetchMongoSkus()
		if err != nil {
			return err
		}
		printObject(skus)
		return nil
	})

}
