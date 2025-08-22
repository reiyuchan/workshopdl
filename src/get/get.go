package get

import (
	"bufio"
	"errors"
	"fmt"
	urlParser "net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func WorkshopGetItem(url string) error {
	efPath, err := os.Executable()
	if err != nil {
		return err
	}

	absPath := filepath.Dir(efPath)

	steamcmd := ""
	appName := ""
	appid := ""
	itemid := ""
	itemName := ""

	switch runtime.GOOS {
	case "windows":
		steamcmd = absPath + "\\steamcmd\\steamcmd.exe"
	default:
		steamcmd = absPath + "/steamcmd/steamcmd.sh"
	}

	_, err = os.Stat(steamcmd)
	if err != nil {
		return err
	}

	parsed, err := urlParser.ParseRequestURI(url)
	if err != nil {
		return err
	}
	itemid = parsed.Query().Get("id")

	c := colly.NewCollector(colly.UserAgent("workshopdl"), colly.AllowedDomains("steamcommunity.com"))

	c.OnHTML("a[data-appid]", func(e *colly.HTMLElement) {
		appid = e.Attr("data-appid")

		fmt.Println("url: " + parsed.String())
		fmt.Println("appid: " + appid)
	})
	c.OnHTML("div.apphub_AppName", func(e *colly.HTMLElement) {
		appName = e.Text
		fmt.Println("appname: " + strings.ToLower(appName))
	})
	c.OnHTML("div.workshopItemTitle", func(e *colly.HTMLElement) {
		itemName = e.Text
		fmt.Println("itemid: " + itemid)
		fmt.Println("itemname: " + strings.ToLower(itemName))
	})

	c.Visit(parsed.String())

	if appid == "" {
		return errors.New("domain not allowed")
	}

	cmd := exec.Command(steamcmd, "+force_install_dir", "installed", "+login", "anonymous", "+workshop_download_item", appid, itemid, "+quit")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "Valve") || strings.Contains(scanner.Text(), "to exit") {
				continue
			}
			if strings.Contains(scanner.Text(), "Redirecting stderr") {
				break
			}
			fmt.Println(scanner.Text())
		}
		wg.Done()
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Fprintln(os.Stderr, "ERR: "+scanner.Text())
		}
		wg.Done()
	}()

	err = cmd.Wait()
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
