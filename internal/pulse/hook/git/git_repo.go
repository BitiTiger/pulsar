/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2021, Lewis Cook <lcook@FreeBSD.org>
 * Copyright (c) 2023, Cameron Himes <cameron.himes@hotmail.com>
 * All rights reserved.
 */
package git

import (
	"strings"
)

type repository struct {
	ID          int    `json:"id,omitempty"`
	NodeID      string `json:"node_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Fullname    string `json:"full_name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r *repository) String() string {
	/*
	* Arbitrarily trim the repository name prefix
	* `freebsd-` since we track the FreeBSD GitHub
	* mirror repositories.
	*
	* `freebsd-ports`, `freebsd-src` and `freebsd-docs`
	* are returned as `ports`, `src`, and `docs`
	* respectively.
	 */
	return strings.TrimPrefix(r.Name, "freebsd-")
}
