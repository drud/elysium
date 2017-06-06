package main

// This simple demonstration app explores a users sites and environments, then gives
// download links to any database or files downloads.

import (
	"github.com/drud/elysium/pkg/elysium"
	"fmt"
	"log"
	"os"
)

func main() {

	session := elysium.NewAuthSession(os.Getenv("TERMINUS_API_TOKEN"))

	SiteList := &elysium.SiteList{}
	err := session.Request("GET", SiteList)

	// Get a list of environments for a given site.
	for _,site := range SiteList.Sites {
		fmt.Printf("\nSite: %s\n", site.Site.Name)
		environmentList := elysium.NewEnvironmentList(site.ID)
		err = session.Request("GET", environmentList)

		for _, envType := range []string{"live", "test", "dev"} {

			// Get a list of all backups for the live.
			env := environmentList.Environments[envType]
			bl := elysium.NewBackupList(site.ID, env.Name)
			err = session.Request("GET", bl)

			// Traverse backups for the site and provide database/files backups.
			if len(bl.Backups) > 0 {
				for _, backup := range bl.Backups {
					if backup.ArchiveType == "database" || backup.ArchiveType == "files" {
						// Get a time-limited backup URL from Pantheon. This requires a POST of the backup type to their API.
						err = session.Request("POST", &backup)
						if err != nil {
							log.Fatal(err)
						}
						// Print the download URL.
						fmt.Printf("\t%s %s %s backup: %s\n", site.Site.Name, envType, backup.ArchiveType, backup.DownloadURL)
					}
				}
			}

		}
	}
}
