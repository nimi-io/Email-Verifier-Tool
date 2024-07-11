package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	log.Printf("domain,hasMx,hasSPF,spfRecord,hasDmarc,dmarcRecord \n")

	for scanner.Scan() {
		domain := strings.ToLower(scanner.Text())
		if domain == "exit" {
			break
		}
		checkDomainName(domain)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input %v\n", err)
	}

}

func checkDomainName(domain string) {

	var hasMx, hasSPF, hasDmarc bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("%v, error\n", domain)
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("%v, error\n", domain)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("%v, error\n", domain)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDmarc = true
			dmarcRecord = record
			break
		}
	}

	log.Printf("domain: %v, \n hasMx: %v, \nhasSPF: %v,\n spfRecord %v,\n hasDmarc: %v,\n dmarcRecord %v, \n", domain, hasMx, hasSPF, spfRecord, hasDmarc, dmarcRecord)

}
