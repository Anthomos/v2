// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui

import "testing"

func TestNormalizeOffset(t *testing.T) {
	testCases := []struct {
		name   string
		total  int
		offset int
		want   int
	}{
		{
			name:   "keeps offset on populated page",
			total:  60,
			offset: 20,
			want:   20,
		},
		{
			name:   "resets offset when it points past last unread page",
			total:  40,
			offset: 40,
			want:   0,
		},
		{
			name:   "keeps zero offset for empty result sets",
			total:  0,
			offset: 0,
			want:   0,
		},
		{
			name:   "keeps non-zero offset for empty result sets",
			total:  0,
			offset: 20,
			want:   20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := normalizeOffset(tc.total, tc.offset); got != tc.want {
				t.Fatalf("normalizeOffset(%d, %d) = %d, want %d", tc.total, tc.offset, got, tc.want)
			}
		})
	}
}
