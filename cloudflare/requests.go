package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strings"

	"github.com/himidori/cf-dynamic-dns/config"
	"github.com/himidori/cf-dynamic-dns/utils"
)

const (
	zonesAPIURL = "https://api.cloudflare.com/client/v4/zones/%s/dns_records"
)

func (cr *CFRequester) Run() {
	for i := 0; i < cr.workers; i++ {
		go cr.processSites()
		cr.wg.Add(1)
	}

	for _, site := range cr.config.Sites {
		cr.workCh <- site
	}
	close(cr.workCh)

	cr.wg.Wait()
}

func (cr *CFRequester) processSites() {
	for site := range cr.workCh {
		cr.logger.Info("processing domain %s", site.Domain)

		records, err := cr.getRecords(&site)
		if err != nil {
			cr.logger.Err("failed to get dns records for domain %s. err: %v", site.Domain, err)
			continue
		}

		currIP, err := utils.GetCurrentIP(cr.config.IPIdentURL)
		if err != nil {
			cr.logger.Err("failed to get current ip. err: %v", err)
			continue
		}

		for _, rec := range records {
			if rec.Content != currIP {
				cr.logger.Info("updating record for domain %s. new ip: %s", rec.DomainName, currIP)

				if err := cr.updateRecord(&site, &rec, currIP); err != nil {
					cr.logger.Err("failed to update record for domain %s. err: %v", rec.DomainName, err)
					continue
				}

				cr.logger.Info("successfuly updated record for domain %s, new ip: %s", rec.DomainName, currIP)
			}
		}
	}
	cr.wg.Done()
}

func (cr *CFRequester) getRecords(site *config.Site) ([]dnsrecord, error) {
	resp, err := cr.makeRequest(
		fmt.Sprintf(zonesAPIURL, site.ZoneID)+"?type=A",
		http.MethodGet,
		site.AuthKey,
		site.AuthEmail,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res := apiResult{}
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}

	rawRecords := []dnsrecord{}
	if err := json.Unmarshal(res.Result, &rawRecords); err != nil {
		return nil, err
	}

	filteredRecords := []dnsrecord{}

	for _, rr := range rawRecords {
		if strings.Contains(rr.DomainName, site.Domain) {
			filteredRecords = append(filteredRecords, rr)
		}
	}

	return filteredRecords, nil
}

func (cr *CFRequester) updateRecord(site *config.Site, record *dnsrecord, content string) error {
	record.Content = content

	bytes, err := json.Marshal(record)
	if err != nil {
		return err
	}

	_, err = cr.makeRequest(
		fmt.Sprintf(zonesAPIURL, site.ZoneID)+"/"+record.ID,
		http.MethodPut,
		site.AuthKey,
		site.AuthEmail,
		nil,
		bytes,
	)

	return err
}
