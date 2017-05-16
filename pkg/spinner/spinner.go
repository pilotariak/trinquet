// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package spinner

import (
	"time"

	gospinner "github.com/briandowns/spinner"
)

var Spinner TextSpinner

type TextSpinner struct {
	spn *gospinner.Spinner
}

func init() {
	Spinner = TextSpinner{
		spn: gospinner.New(gospinner.CharSets[9], 60*time.Millisecond),
	}
}

func (ts *TextSpinner) Start(suffix string) {
	ts.spn.Suffix = " " + suffix
	ts.spn.Start()
}

func (ts *TextSpinner) Stop() {
	ts.spn.Stop()
}
