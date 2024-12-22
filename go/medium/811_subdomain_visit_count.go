// https://leetcode.com/problems/subdomain-visit-count/
package medium

import (
	"fmt"
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) []string {
	counts := make(map[string]int)
	for _, cpdomain := range cpdomains {
		count, subDomains := parseCPDomain(cpdomain)
		for _, subDomain := range subDomains {
			counts[subDomain] += count
		}
	}

	result := make([]string, 0)
	for domain, count := range counts {
		result = append(result, fmt.Sprintf("%d %s", count, domain))
	}
	return result
}

func parseCPDomain(cpdomain string) (int, []string) {
	parts := strings.Split(cpdomain, " ")
	count, _ := strconv.Atoi(parts[0])
	domain := parts[1]

	subDomains := make([]string, 0)
	domainParts := strings.Split(domain, ".")
	for i := range domainParts {
		subDomains = append(subDomains, strings.Join(domainParts[i:], "."))
	}

	return count, subDomains
}
