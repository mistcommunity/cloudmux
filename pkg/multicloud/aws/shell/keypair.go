package shell

import (
	"fmt"

	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/aws"
)

func init() {
	type KeyPairListOptions struct {
		Name   string
		Finger string
	}
	shellutils.R(&KeyPairListOptions{}, "keypair-list", "List keypairs", func(cli *aws.SRegion, args *KeyPairListOptions) error {
		keypairs, err := cli.GetKeypairs(args.Finger, args.Name)
		if err != nil {
			return err
		}
		printList(keypairs, 0, 0, 0, []string{})
		return nil
	})

	type KeyPairImportOptions struct {
		NAME   string `help:"Name of new keypair"`
		PUBKEY string `help:"Public key string"`
	}
	shellutils.R(&KeyPairImportOptions{}, "keypair-import", "Import a keypair", func(cli *aws.SRegion, args *KeyPairImportOptions) error {
		keypair, err := cli.ImportKeypair(args.NAME, args.PUBKEY)
		if err != nil {
			return err
		}
		printObject(keypair)
		return nil
	})

	type KeyPairSyncOptions struct {
		PUBKEY string `help:"Public key string"`
	}
	shellutils.R(&KeyPairSyncOptions{}, "keypair-sync", "Sync a keypair", func(cli *aws.SRegion, args *KeyPairSyncOptions) error {
		key, err := cli.SyncKeypair(args.PUBKEY)
		if err != nil {
			return err
		}
		fmt.Println(key)
		return nil
	})

}
