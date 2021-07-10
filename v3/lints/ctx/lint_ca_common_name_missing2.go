package ctx

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type caCommonNameMissing struct {
	BeerHall string
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ca_common_name_missing2",
		Description:   "CA Certificates common name MUST be included.",
		Citation:      "BRs: 7.1.4.3.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABV148Date,
		Lint:          NewCaCommonNameMissing,
	})
}

func (l *caCommonNameMissing) Configure() interface{} {
	return l
}

func NewCaCommonNameMissing() lint.LintInterface {
	return &caCommonNameMissing{}
}

func (l *caCommonNameMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c)
}

func (l *caCommonNameMissing) Execute(c *x509.Certificate) *lint.LintResult {
	fmt.Printf("BeerHall name is: '%s'\n", l.BeerHall)
	if c.Subject.CommonName == "" {
		return &lint.LintResult{Status: lint.Error, Details: l.BeerHall}
	} else {
		return &lint.LintResult{Status: lint.Pass, Details: l.BeerHall}
	}
}
