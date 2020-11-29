package rfc

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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
	"testing"

	"github.com/cpu/yylint/v2/lint"
	"github.com/cpu/yylint/v2/test"
)

func TestSubjectOrganizationalUnitNameLengthGood(t *testing.T) {
	inputPath := "subjectOrganizationalUnitNameLengthGood.pem"
	expected := lint.Pass
	out := test.TestLint("e_subject_organizational_unit_name_max_length", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectOrganzationalUnitNameLong(t *testing.T) {
	inputPath := "subjectOrganizationalUnitNameLong.pem"
	expected := lint.Error
	out := test.TestLint("e_subject_organizational_unit_name_max_length", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
